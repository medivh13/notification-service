package usecases

import (
	mailUC "notification/src/app/usecase/mail"
)

type AllUseCases struct {
	MailUseCase mailUC.MailUsecaseInterface
}
