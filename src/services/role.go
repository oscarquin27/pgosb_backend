package services

import (
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
	"fdms/src/utils/results"
)

type RoleService interface {
	abstract_handler.AbstractCRUDService[models.Role]
	GetSchema(id int64) *results.ResultWithValue[string]
}
