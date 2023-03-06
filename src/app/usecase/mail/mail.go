package mailusecase

import (
	"bytes"
	"html/template"
	"log"
	dto "notification/src/app/dtos/email"
	notifConst "notification/src/infra/constant"
	emailSrv "notification/src/infra/services/mail"

	"github.com/alitto/pond"
)

var groupCount = 0

type MailUsecaseInterface interface {
	MailNotif(data *dto.EmailReqDTO) error
}

type mailUseCase struct {
	EmailServices    emailSrv.EmailServicesInt
	Pondnotification *pond.WorkerPool
}

func NewMailUseCase(emailServices emailSrv.EmailServicesInt, pondnotification *pond.WorkerPool) *mailUseCase {
	return &mailUseCase{
		EmailServices:    emailServices,
		Pondnotification: pondnotification,
	}
}

func (uc *mailUseCase) MailNotif(data *dto.EmailReqDTO) error {

	group := uc.Pondnotification.Group()
	group.Submit(func() {

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
	})
	group.Wait()
	return nil
}
