package stuff

import (
	"reflect"
	"serviceX/src/model"
	"strings"

	"gorm.io/gorm"
)

type Stuff struct {
	controller model.ControllerI `json:"-" example:"1"`
	model.BaseSQL
	Field1 int    `json:"field1" example:"1" binding:"max=200"`
	Name   string `json:"name" example:"peter"`
}

// GetName Return the struct/model name
func (m *Stuff) GetName() string {
	//We can remove reflect and juste set a constant
	return strings.ToLower(reflect.TypeOf(m).Elem().Name()) // using the struct name as a collection name
	//return "stuff"
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
func (m *Stuff) OnCreate() (status int, err error) { return }
func (m *Stuff) OnUpdate() (status int, err error) { return }
func (m *Stuff) OnDelete() (status int, err error) { return }

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
