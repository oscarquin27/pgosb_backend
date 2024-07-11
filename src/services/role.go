package services

import "fdms/src/models"

type RoleService interface {
	Get(id int64) (*models.Role, error)
	GetSchema(id int64) (*string, error)
	GetAll() ([]models.Role, error)
	Create(user *models.Role) error
	Update(user *models.Role) error
	Delete(id int64) error
}
