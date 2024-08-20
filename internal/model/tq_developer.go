package model

import "gorm.io/gorm"

type TqDeveloper struct {
	gorm.Model
}

func (m *TqDeveloper) TableName() string {
	return "tq_developer"
}
