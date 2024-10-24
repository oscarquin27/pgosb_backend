package services

import "fdms/src/models"

type MissionPersonService interface {
	Get(id int) (*models.MissionPerson, error)
	GetAll() ([]models.MissionPerson, error)
	GetMissionId(id int) ([]models.MissionPerson, error)
	Create(user *models.MissionPerson) error
	Update(user *models.MissionPerson) error
	Delete(id int) error
}
