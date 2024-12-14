package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const TaskSendVerifyEmail = "task:send_verify_email"

type PayloadSendVerifyEmail struct {
	ID uuid.UUID `json:"id"`
}

func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("ペイロードのマーシャリングに失敗しました: %w", err)
	}
	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("エンキューに失敗しました: %w", err)
	}

	log.Info().Str("type", task.Type()).
		Bytes("payload", task.Payload()).
		Str("queue", info.Queue).
		Int("max_retry", info.MaxRetry).
		Msg("タスクをエンキュー")
	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("ペイロードのアンマーシャリングに失敗: %w", asynq.SkipRetry)
	}

	verifyEmail, err := processor.store.GetVerifyEmail(ctx, payload.ID)
	if err != nil {
		return fmt.Errorf("仮登録アカウント取得に失敗しました: %w", err)
	}
	verifyURL := fmt.Sprintf("https://www.iput-kernel.com/verify_email?email_id=%s&secret_code=%s",
		verifyEmail.ID, verifyEmail.SecretCode)

	err = processor.mailer.SendConfirmationMail(verifyEmail.Name, verifyEmail.Email, verifyURL)

	if err != nil {
		return fmt.Errorf("認証メールを送信できませんでした: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("email", verifyEmail.Email).Msg("タスクを処理")
	return nil
}
