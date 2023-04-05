package mail

import (
	"testing"

	"github.com/lushenle/mundilfari/util"
	"github.com/stretchr/testify/require"
)

func TestMailSender(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewMailgunSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	<p>This is a test message from <a href="https://shenle.lu">Mundilfari</a></p>
	`
	to := []string{"lushenle@gmail.com"}
	cc := []string{"zhinouk@gmail.com"}
	attachFiles := []string{"../main.go"}
	err = sender.Sender(subject, content, to, cc, nil, attachFiles)
	require.NoError(t, err)
}
