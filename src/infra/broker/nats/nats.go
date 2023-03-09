package nats

import (
	"log"
	"notification/src/infra/config"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
) /*
 * Author      : Jody (github.com/medivh13)
 * Modifier    :
 * Domain      : notification
 */

type Nats struct {
	Status bool
	Conn   *nats.Conn
}

func NewNats(conf config.NatsConf, logger *logrus.Logger) *Nats {
	var Nats = new(Nats)

	if conf.NatsStatus == "1" {
		Nats.Status = true
	}

	if Nats.Status {

		var err error
		Nats.Conn, err = nats.Connect(conf.NatsHost)

		if err != nil {
			logger.Printf("error connecting NATS. %s\n", err.Error())
		}
		log.Println("connected to:", conf.NatsHost)
	}

	return Nats
}
