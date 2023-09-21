package model

import (
	"gorm.io/gorm"
)

type Stuff struct {
	BaseModel
	Field1 int `json:"field1" example:"1"`
}

func (m *Stuff) BeforeValidation() {
}

func (m *Stuff) AfterValidation() {
}

func (m *Stuff) CreateValidation() (bool, map[string]string) {
	return false, nil
}

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
