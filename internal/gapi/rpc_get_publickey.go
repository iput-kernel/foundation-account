package gapi

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/iput-kernel/foundation-account/internal/pb"
)

func (server *Server) GetPublicKey(ctx context.Context, req *pb.Empty) (*pb.GetPublicKeyResponse, error) {
	x509EncodedPub, err := x509.MarshalPKIXPublicKey(server.publicKey)
	if err != nil {
		fmt.Println("Error encoding public key:", err)
		return nil, err
	}

	// PEM形式にエンコード
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509EncodedPub,
	})

	rsp := &pb.GetPublicKeyResponse{
		PublicKey: string(pemEncodedPub),
	}
	return rsp, nil
}
