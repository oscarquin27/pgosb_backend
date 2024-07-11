package services

import "fdms/src/models"

type UserService interface {
	Get(id int64) (*models.User, error)
	GetAll() ([]models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id int64) error
}
