package mail

import (
	"testing"

	"github.com/iput-kernel/foundation-account/internal/config"
	"github.com/stretchr/testify/require"
)

const TestAddress = "example@email.com"

func TestSendEmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config, err := config.LoadConfig("../../..")
	require.NoError(t, err)

	sender := NewEmailSender(config.EmailSender.Name, config.EmailSender.Address, config.EmailSender.Password)

	subject := "テストメール"
	content := `
	<h1>Hello world</h1>
	<p>これは <a href="http://www.iput-kernel.com">IPUT-Kernel</a>からのテストメッセージです</p>
	`
	to := []string{TestAddress}

	err = sender.SendEmail(subject, content, to, nil, nil, nil)
	require.NoError(t, err)
}

func TestSendConfirmationEmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config, err := config.LoadConfig("../../..")
	require.NoError(t, err)
	sender := NewSendConfirmationMail(config.EmailSender.Name, config.EmailSender.Address, config.EmailSender.Password)
	err = sender.SendConfirmationMail("テストメールクライアントさん", TestAddress, "https://example.com")
	require.NoError(t, err)
}
