package server

import (
	"context"
	"example-nunu/internal/service"
	"example-nunu/pkg/log"
)

type Job struct {
	log    *log.Logger
	tsApps service.TqAppService
}

func NewJob(
	log *log.Logger,
) *Job {
	return &Job{
		log: log,
	}
}
func (j *Job) Start(ctx context.Context) error {
	// eg: kafka consumer
	return nil
}
func (j *Job) Stop(ctx context.Context) error {
	return nil
}
