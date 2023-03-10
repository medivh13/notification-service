package main

import (
	"context"
	"fmt"

	usecases "notification/src/app/usecase"
	mailUc "notification/src/app/usecase/mail"

	"notification/src/infra/config"

	"notification/src/interface/rest"

	emailSrv "notification/src/infra/services/mail"

	ms_log "notification/src/infra/log"

	"github.com/alitto/pond"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// init context
	ctx := context.Background()

	// read the server environment variables
	conf := config.Make()

	// check is in production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// logger setup
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	fmt.Printf("%+v", conf.SMTP)
	emailIntegration := emailSrv.NewEmailService(conf.SMTP)

	pool := pond.New(10, 1000)
	defer pool.StopAndWait()

	// Create a task group
	// HTTP Handler
	// the server already implements a graceful shutdown.

	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		usecases.AllUseCases{
			MailUseCase: mailUc.NewMailUseCase(emailIntegration, pool),
		},
	)
	if err != nil {
		panic(err)
	}
	httpServer.Start(ctx)

}
