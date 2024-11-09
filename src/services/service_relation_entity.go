package services

import (
	"context"
	"fdms/src/models"
	"fdms/src/utils/results"
)

type ServiceRelationEntityService interface {
	Create(ctx context.Context, s *models.ServiceRelationEntity) *results.ResultWithValue[models.ServiceRelationEntity]
	Update(ctx context.Context, s *models.ServiceRelationEntity) *results.ResultWithValue[models.ServiceRelationEntity]
}
