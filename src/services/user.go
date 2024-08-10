package services

import (
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
	"fdms/src/utils/results"
)

type UserService interface {
	abstract_handler.AbstractCRUDService[models.User]
	GetAllSimple() *results.ResultWithValue[[]models.UserSimple]
}
