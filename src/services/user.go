package services

import (
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
)

type UserService interface {
	abstract_handler.AbstractCRUDService[models.User]
}
