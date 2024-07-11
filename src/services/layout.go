package services

import "fdms/src/models"

type LayoutService interface {
	Get(entity string) ([]models.Layout, error)
}
