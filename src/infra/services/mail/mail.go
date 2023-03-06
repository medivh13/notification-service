package mail

/*
 * Author      : Jody (github.com/medivh13)
 * Modifier    :
 * Domain      : notification
 */
import (
	"log"
	"strings"

	"notification/src/infra/config"

	"github.com/go-mail/mail"
)

type EmailServicesInt interface {
	SendEmail(sender, title, message string, receiver []string) error
}
type emailService struct {
	mailAuth config.SMTPConf
}

func NewEmailService(smtp config.SMTPConf) EmailServicesInt {
	return &emailService{
		mailAuth: smtp,
	}
}

func (s *emailService) SendEmail(sender, title, message string, receiver []string) error {

	toHeader := strings.Join(receiver, ",")
	m := mail.NewMessage()

	m.SetHeader("From", sender)

	m.SetHeader("To", toHeader)

	m.SetHeader("Subject", title)

	m.SetBody("text/html", message)

	d := mail.NewDialer(s.mailAuth.Host, s.mailAuth.Port, s.mailAuth.UserName, s.mailAuth.Password)

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		return err

	}
	return nil
}
