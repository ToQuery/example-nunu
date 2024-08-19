package model

import "gorm.io/gorm"

type App struct {
	gorm.Model
}

func (m *App) TableName() string {
	return "app"
}
