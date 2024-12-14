package mail

import (
	"bytes"
	"path/filepath"
	"text/template"
)

type ConfirmationEmailSender struct {
	sender EmailSender
}

func NewSendConfirmationMail(senderName, fromEmailAddress, fromEmailPassword string) *ConfirmationEmailSender {
	sender := NewEmailSender(senderName, fromEmailAddress, fromEmailPassword)
	return &ConfirmationEmailSender{
		sender: sender,
	}
}

func (c *ConfirmationEmailSender) SendConfirmationMail(username string, to string, link string) error {
	templatesDir := filepath.Join("internal", "application", "mail", "templates", "confirmation.html")

	// HTMLをロード
	tmpl, err := template.ParseFiles(templatesDir)
	if err != nil {
		return err
	}

	// Prepare the data for the template
	data := struct {
		Username    string
		Link        string
		AcademyName string
	}{
		Username:    username,
		Link:        link,
		AcademyName: "東京国際工科専門職大学",
	}

	// Execute the template with the data
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return err
	}

	err = c.sender.SendEmail(
		"登録確認リンク",     // メールタイトル
		body.String(), // 内容
		[]string{to},  // 宛先
		nil,           // cc
		nil,           // bcc
		nil,           // 添付ファイル
	)
	if err != nil {
		return err
	}

	return nil
}
