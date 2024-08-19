package service

import (
	"context"
	"example-nunu/internal/model"
	"example-nunu/internal/repository"
)

type AppService interface {
	GetApp(ctx context.Context, id int64) (*model.App, error)
}

func NewAppService(
	service *Service,
	appRepository repository.AppRepository,
) AppService {
	return &appService{
		Service:       service,
		appRepository: appRepository,
	}
}

type appService struct {
	*Service
	appRepository repository.AppRepository
}

func (s *appService) GetApp(ctx context.Context, id int64) (*model.App, error) {
	return s.appRepository.GetApp(ctx, id)
}
