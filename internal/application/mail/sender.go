package mail

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	//　認証サーバー 基本的にServerAddressと同じホスト名にしておけばいい。
	smtpAuthAddress   = "mail.iput-kernel.com"
	smtpServerAddress = "mail.iput-kernel.com:465"
)

type EmailSender interface {
	SendEmail(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		attachFiles []string,
	) error
}

type Email struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewEmailSender(name string, fromEmailAddress string, fromEmailPassword string) EmailSender {
	return &Email{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *Email) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("添付ファイル取得に失敗 %s: %w", f, err)
		}
	}

	// TLS設定
	tlsConfig := &tls.Config{
		ServerName: smtpAuthAddress,
	}
	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, smtpAuthAddress)

	//メール送信
	if err := e.SendWithTLS(smtpServerAddress, smtpAuth, tlsConfig); err != nil {
		return fmt.Errorf("メール送信に失敗: %w", err)
	}

	return nil
}
