package model

import "gorm.io/gorm"

type TqApp struct {
	gorm.Model
}

func (m *TqApp) TableName() string {
	return "tq_app"
}
