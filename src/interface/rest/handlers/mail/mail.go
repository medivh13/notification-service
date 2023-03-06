package brand_handlers

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : notification
 */

import (
	"encoding/json"
	"net/http"

	dto "notification/src/app/dtos/email"
	usecases "notification/src/app/usecase/mail"
	common_error "notification/src/infra/errors"
	"notification/src/interface/rest/response"
)

type MailHandlerInterface interface {
	Createmail(w http.ResponseWriter, r *http.Request)
}

type mailHandler struct {
	response response.IResponseClient
	usecase  usecases.MailUsecaseInterface
}

func NewMaildHandler(r response.IResponseClient, h usecases.MailUsecaseInterface) MailHandlerInterface {
	return &mailHandler{
		response: r,
		usecase:  h,
	}
}

func (h *mailHandler) Createmail(w http.ResponseWriter, r *http.Request) {

	postDTO := dto.EmailReqDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	err = postDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	err = h.usecase.MailNotif(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Send Notif",
		nil,
		nil,
	)
}
