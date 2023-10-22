package stuff

import (
	controller "serviceX/src/api"
	"serviceX/src/model"

	"gorm.io/gorm"
)

type Stuff struct {
	*controller.Controller
	model.BaseModel
	Field1 int    `json:"field1" example:"1"`
	Name   string `json:"name" example:"peter"`
}

func (m *Stuff) BeforeValidation() {
	m.BaseModel.BeforeValidation()
}

func (m *Stuff) AfterValidation() {
	m.BaseModel.AfterValidation()
}

func (m *Stuff) CreateValidation() (bool, map[string]string) {
	return true, nil
}

// Hooks for Gorm
func (m *Stuff) BeforeCreate(tx *gorm.DB) error {
	return nil
}

func (m *Stuff) AfterDelete(tx *gorm.DB) error {
	return nil
}

func (m *Stuff) AfterUpdate(tx *gorm.DB) error {
	return nil
}

func (m *Stuff) AfterCreate(tx *gorm.DB) error {
	return nil
}
