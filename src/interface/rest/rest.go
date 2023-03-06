package rest

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	usecases "notification/src/app/usecase"
	"notification/src/infra/config"

	"notification/src/interface/rest/handlers"
	"notification/src/interface/rest/response"
	"notification/src/interface/rest/route"

	mailHandler "notification/src/interface/rest/handlers/mail"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// HttpServer holds the dependencies for a HTTP server.
type HttpServer struct {
	*http.Server
	logger *logrus.Logger
}

// New creates and configures a server serving all application routes.
// The server implements a graceful shutdown.
// chi.Mux is used for registering some convenient middlewares and easy configuration of
// routes using different http verbs.
func New(
	conf config.HttpConf,
	isProd bool,
	logger *logrus.Logger,
	useCases usecases.AllUseCases,
) (*HttpServer, error) {
	// wrap all the routes
	routeHandler := makeRoute(conf.XRequestID, conf.Timeout, isProd, logger, useCases)

	// http service
	srv := http.Server{
		Addr:    ":" + conf.Port,
		Handler: routeHandler,
	}

	return &HttpServer{&srv, logger}, nil
}

// makeRoute register routes
func makeRoute(
	xRequestID string,
	timeout int,
	isProd bool,
	logger *logrus.Logger,
	useCases usecases.AllUseCases,
) *chi.Mux {

	r := chi.NewRouter()

	// apply common middleware here ...

	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	// logging middleware
	// if !isProd {
	// 	r.Use(middleware.Logger)
	// }
	r.Use(middleware.Logger)

	// if isProd {
	// 	// relic middleware
	// 	r.Use(ms_middleware.InitNewRelicHandler(newRelic.AppName, newRelic.License))
	// }

	// x-request-id middleware
	if len(xRequestID) <= 0 {
		logger.Fatalf("invalid x-request-id")
	}

	// timeout middleware
	if timeout <= 0 {
		logger.Fatalf("invalid http timeout")
	}

	// instantiate the handlers here ...

	respClient := response.NewResponseClient()
	hh := handlers.NewHealthHandler(respClient)
	mh := mailHandler.NewMaildHandler(respClient, useCases.MailUseCase)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/check", route.HealthRouter(hh))
		r.Mount("/mail", route.MailAppRouter(mh))

	})

	return r
}

// Start runs ListenAndServe on the http.Server with graceful shutdown
func (srv *HttpServer) Start(ctx context.Context) {
	// run HTTP service
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			srv.logger.Fatal(err)
		}
	}()

	// ready to serve
	srv.logger.Info("listen on", srv.Addr)

	srv.gracefulShutdown(ctx)
}

func (srv *HttpServer) gracefulShutdown(ctx context.Context) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// wait forever until 1 of signals above are received
	<-quit
	srv.logger.Warnf("got signal: %v, shutting down server ...", quit)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		srv.logger.Error(err)
	}

	srv.logger.Println("server exiting")
}
