package services

import "fdms/src/models"

type PersonService interface {
	Get(id int) (*models.User, error)
	GetAll() ([]models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id int64) error
}
