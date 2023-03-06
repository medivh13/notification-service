package route

import (
	"net/http"

	mailHandler "notification/src/interface/rest/handlers/mail"

	"github.com/go-chi/chi/v5"
)

func MailRouter(h mailHandler.MailHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/", h.Createmail)
	return r
}
