package main

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hibiken/asynq"
	"github.com/iput-kernel/foundation-account/internal/application/mail"
	"github.com/iput-kernel/foundation-account/internal/config"
	"github.com/iput-kernel/foundation-account/internal/gapi/method"
	"github.com/iput-kernel/foundation-account/internal/gapi/service"
	"github.com/iput-kernel/foundation-account/internal/infra/db/repository"
	"github.com/iput-kernel/foundation-account/internal/infra/worker"
	accountv1 "github.com/iput-kernel/foundation-account/internal/pb/account/service/v1"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("設定が読み込めませんでした。")
	}

	if config.Environment == "DEV" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	connPool, err := pgxpool.New(ctx, config.DSN())
	if err != nil {
		log.Fatal().Err(err).Msg("データベースに接続できませんでした")
	}

	store := repository.NewDAO(connPool)

	redisOpt := asynq.RedisClientOpt{
		Addr:     net.JoinHostPort(config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)

	waitGroup, ctx := errgroup.WithContext(ctx)

	runTaskProcessor(ctx, waitGroup, config, redisOpt, store)
	runGrpcServer(ctx, waitGroup, config, store, taskDistributor)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("wait groupでエラーが発生")
	}
}

func runTaskProcessor(
	ctx context.Context,
	waitGroup *errgroup.Group,
	config config.Config,
	redisOpt asynq.RedisClientOpt,
	store repository.DAO,
) {
	mailer := *mail.NewSendConfirmationMail(config.EmailSender.Name, config.EmailSender.Address, config.EmailSender.Password)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)

	log.Info().Msg("タスクプロセッサーを起動")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("タスクプロセッサーの起動に失敗")
	}

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("タスクプロセッサーをgraceful shutdown")

		taskProcessor.Shutdown()
		log.Info().Msg("タスクプロセッサーが終了しました。")

		return nil
	})
}

func runGrpcServer(
	ctx context.Context,
	waitGroup *errgroup.Group,
	config config.Config,
	store repository.DAO,
	taskDistributor worker.TaskDistributor,
) {
	server, err := service.NewServer(config, store, taskDistributor)
	method := method.NewMethod(server)
	if err != nil {
		log.Fatal().Err(err).Msg("serverが作成できません")
	}

	gprcLogger := grpc.UnaryInterceptor(service.GrpcLogger)
	grpcServer := grpc.NewServer(gprcLogger)
	accountv1.RegisterAccountServiceServer(grpcServer, method)
	reflection.Register(grpcServer)

	GRPCServerAddress := net.JoinHostPort(config.Server.Host, config.Server.GRPCPort)
	listener, err := net.Listen("tcp", GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("listenerの作成に失敗")
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("gRPCサーバー起動: %s", listener.Addr().String())

		err = grpcServer.Serve(listener)
		if err != nil {
			if errors.Is(err, grpc.ErrServerStopped) {
				return nil
			}
			log.Error().Err(err).Msg("gRPC serverがserveできません")
			return err
		}

		return nil
	})

	// graceful shutdown
	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("gRPCサーバーをgraceful shutdown")

		grpcServer.GracefulStop()
		log.Info().Msg("gRPCサーバーが終了")

		return nil
	})
}
