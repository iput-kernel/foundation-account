package method

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	accountv1 "github.com/iput-kernel/foundation-account/internal/pb/account/auth/v1"
)

func (server *Method) GetPublicKey(ctx context.Context, req *accountv1.GetPublicKeyRequest) (*accountv1.GetPublicKeyResponse, error) {
	x509EncodedPub, err := x509.MarshalPKIXPublicKey(server.PublicKey)
	if err != nil {
		fmt.Println("Error encoding public key:", err)
		return nil, err
	}

	// PEM形式にエンコード
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509EncodedPub,
	})

	rsp := &accountv1.GetPublicKeyResponse{
		PublicKey: string(pemEncodedPub),
	}
	return rsp, nil
}
