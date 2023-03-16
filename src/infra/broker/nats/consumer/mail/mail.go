package mail

import (
	"encoding/json"
	"log"

	dto "notification/src/app/dtos/email"
	natsBroker "notification/src/infra/broker/nats"
	notifConst "notification/src/infra/constant"

	useCase "notification/src/app/usecase/mail"

	"github.com/nats-io/nats.go"
)

const (
	FcmMaxRecipient    = 1000
	FcmMaxRecipientSDK = 500
	PriorityHigh       = "high"
)

type NotifWorkerInterface interface {
	InitNats()
}

type SentNotifWorkerImpl struct {
	nats     *natsBroker.Nats
	subjects string
	queues   string
	UseCase  useCase.MailUsecaseInterface
}

func NewNotifWorker(
	Nats *natsBroker.Nats, useCase useCase.MailUsecaseInterface,
) NotifWorkerInterface {

	subjects := notifConst.NOTIF_SENT

	queues := notifConst.NOTIF_SENT_QUEUE

	sentNotifWorkerImpl := &SentNotifWorkerImpl{
		nats:     Nats,
		subjects: subjects,
		queues:   queues,
		UseCase:  useCase,
	}

	if Nats.Status {
		sentNotifWorkerImpl.InitNats()
	}

	return sentNotifWorkerImpl
}

func (p *SentNotifWorkerImpl) InitNats() {

	go eventNotificationWorker(p)

}

func eventNotificationWorker(t *SentNotifWorkerImpl) {

	_, err := t.nats.Conn.QueueSubscribe(t.subjects, t.queues, func(msg *nats.Msg) {

		dataConsume := dto.EmailReqDTO{}
		err := json.Unmarshal(msg.Data, &dataConsume)
		if err != nil {
			log.Printf("%+v", err)

			return
		}

		err = t.UseCase.MailSent(&dataConsume)
		if err != nil {
			log.Printf("%+v", err)

			return
		}
	})

	if err != nil {
		log.Fatal(err)
	}
	t.nats.Conn.Flush()
	if err := t.nats.Conn.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]", t.subjects)

}
