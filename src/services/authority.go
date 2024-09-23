package services

import (
	"fdms/src/infrastructure/abstract_handler"
	"fdms/src/models"
)

type AuthorityService interface {
	abstract_handler.AbstractCRUDService[models.Authority]
}
