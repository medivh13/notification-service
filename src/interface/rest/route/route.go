package route

import (
	mailHandler "notification/src/interface/rest/handlers/mail"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func MailAppRouter(mh mailHandler.MailHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Mount("/", MailRouter(mh))

	return r
}
