package handlers

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : notification
 */
import (
	"net/http"

	"notification/src/interface/rest/response"
)

type IHealthHandler interface {
	Ping(w http.ResponseWriter, r *http.Request)
	HealthCheck(w http.ResponseWriter, r *http.Request)
}

type healthHandler struct {
	response response.IResponseClient
}

func NewHealthHandler(r response.IResponseClient) IHealthHandler {
	return &healthHandler{
		response: r,
	}
}

func (h *healthHandler) Ping(w http.ResponseWriter, r *http.Request) {
	h.response.JSON(w, "Pong", nil, nil)
}

func (h *healthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {

	payload := response.HealthCheckMessage{
		ServiceName: "JENKINS_SERVICENAME",
		Version:     "JENKINS_GIT_TAG",
		CommitId:    "JENKINS_GIT_COMMIT",
		UpdatedAt:   "JENKINS_TIME",
		Status:      "pass",
	}

	h.response.JSON(w, "OK", payload, nil)
}
