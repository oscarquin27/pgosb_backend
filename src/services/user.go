package services

import (
	"fdms/src/models"
	"fdms/src/utils/results"
)

type UserService interface {
	Get(id int64) (*models.User, error)
	GetAll() ([]models.User, error)
	Create(user *models.User) *results.Result
	Update(user *models.User) error
	Delete(id int64) error
}
