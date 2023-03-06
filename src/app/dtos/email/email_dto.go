package email_dto

/*
 * Author      : Jody (github.com/medivh13)
 * Modifier    :
 * Domain      : notification-service
 */
import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type EmailDTOInterface interface {
	Validate() error
}

type EmailReqDTO struct {
	Recipients []*RecipientsDTO `json:"recipients"`
	Title      string           `json:"title"`
	Message    string           `json:"message"`
	Sender     string           `json:"sender"`
}

type RecipientsDTO struct {
	Email string `json:"email"`
}

func NewEmailDTO(datarecipients []string, title, message, sender string) *EmailReqDTO {
	var recipient RecipientsDTO
	var recipients []*RecipientsDTO

	for _, val := range datarecipients {
		recipient.Email = val

		recipients = append(recipients, &recipient)
	}
	return &EmailReqDTO{
		Recipients: recipients,
		Title:      title,
		Message:    message,
		Sender:     sender,
	}
}

func (dto *EmailReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Recipients, validation.Required),
		validation.Field(&dto.Message, validation.Required),
		validation.Field(&dto.Sender, validation.Required, is.Email),
	); err != nil {
		return err
	}
	return nil
}
