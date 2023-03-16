package mailusecase

import (
	"bytes"
	"encoding/json"
	"log"
	dto "notification/src/app/dtos/email"
	natsPublisher "notification/src/infra/broker/nats/publisher"
	notifConst "notification/src/infra/constant"
	mailService "notification/src/infra/services/mail"
	"text/template"
)

var groupCount = 0

type MailUsecaseInterface interface {
	MailNotif(data *dto.EmailReqDTO) error
	MailSent(data *dto.EmailReqDTO) error
}

type mailUseCase struct {
	MailPublisher natsPublisher.PublisherInterface
	EmailServices mailService.EmailServicesInt
}

func NewMailUseCase(mailsServ mailService.EmailServicesInt, mailPublisher natsPublisher.PublisherInterface) *mailUseCase {
	return &mailUseCase{
		MailPublisher: mailPublisher,
		EmailServices: mailsServ,
	}
}

func (uc *mailUseCase) MailNotif(data *dto.EmailReqDTO) error {

	newData, _ := json.Marshal(data)
	err := uc.MailPublisher.Nats(newData, notifConst.NOTIF_SENT)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (uc *mailUseCase) MailSent(data *dto.EmailReqDTO) error {

	var receiver []string
	for _, val := range data.Recipients {
		receiver = append(receiver, val.Email)
	}
	tmpl, err := template.New("webpage").Parse(notifConst.MailTemplate)

	if err != nil {
		log.Println(err)

	}
	buf := new(bytes.Buffer)
	templateData := struct {
		Message string
	}{
		Message: data.Message,
	}

	err = tmpl.Execute(buf, templateData)
	if err != nil {
		log.Println(err)

	}
	err = uc.EmailServices.SendEmail(data.Sender, data.Title, buf.String(), receiver)
	if err != nil {
		log.Println(err)
	}

	return nil
}
