package mail

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddress   = "smtp.mailgun.org"
	smtpServerAddress = "smtp.mailgun.org:25"
)

type EmailSender interface {
	Sender(subject, content string, to, cc, bcc, attachFiles []string) error
}

type MailgunSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewMailgunSender(name, fromEmailAddress, fromEmailPassword string) EmailSender {
	return &MailgunSender{name: name, fromEmailAddress: fromEmailAddress, fromEmailPassword: fromEmailPassword}
}

func (g *MailgunSender) Sender(subject, content string, to, cc, bcc, attachFiles []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", g.name, g.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	for _, file := range attachFiles {
		_, err := e.AttachFile(file)
		if err != nil {
			return fmt.Errorf("failed to attach file %s: %w", file, err)
		}
	}

	smtpAuth := smtp.PlainAuth("", g.fromEmailAddress, g.fromEmailPassword, smtpAuthAddress)
	return e.Send(smtpServerAddress, smtpAuth)
}
