package stuff

import (
	"serviceX/src/model"

	"gorm.io/gorm"
)

type Stuff struct {
	controller model.ControllerI `json:"-" example:"1"`
	model.BaseSQL
	Field1 int    `json:"field1" example:"1 binding:"max=200"`
	Name   string `json:"name" example:"peter"`
}

func (m *Stuff) SetController(c model.ControllerI) {
	m.controller = c
}
func (m *Stuff) BeforeValidation() {
	m.BaseSQL.BeforeValidation()
}

func (m *Stuff) AfterValidation() {
	m.BaseSQL.AfterValidation()
}

func (m *Stuff) CreateValidation() (bool, map[string]string) {
	return true, nil
}
func (m *Stuff) OnCreate() {}
func (m *Stuff) OnUpdate() {}
func (m *Stuff) OnDelete() {}

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
