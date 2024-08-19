package repository

import (
	"context"
	"example-nunu/internal/model"
)

type AppRepository interface {
	GetApp(ctx context.Context, id int64) (*model.App, error)
}

func NewAppRepository(
	repository *Repository,
) AppRepository {
	return &appRepository{
		Repository: repository,
	}
}

type appRepository struct {
	*Repository
}

func (r *appRepository) GetApp(ctx context.Context, id int64) (*model.App, error) {
	var app model.App

	return &app, nil
}
