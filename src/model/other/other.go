package other

import (
	"reflect"
	"serviceX/src/model"
	"strings"
	"time"
)

type Other struct {
	model.BaseNoSQL `bson:"inline"`
	controller      model.ControllerI `json:"-"`
	Field1          int               `json:"field1" example:"1" binding:"max=200"`
	Name            string            `json:"name" example:"peter"`
}

// GetName Return the struct/model name
func (m *Other) GetName() string {
	//We can remove reflect and juste set a constant
	return strings.ToLower(reflect.TypeOf(m).Elem().Name()) // using the struct name as a collection name
	//return "other"
}

func (m *Other) SetController(c model.ControllerI) {
	m.controller = c
}
func (m *Other) BeforeValidation() {
	m.BaseNoSQL.BeforeValidation()
}

func (m *Other) AfterValidation() {
	m.BaseNoSQL.AfterValidation()
}

func (m *Other) CreateValidation() (bool, map[string]string) {
	return true, nil
}
func (m *Other) OnCreate() (status int, err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}
func (m *Other) OnUpdate() (status int, err error) {
	m.UpdatedAt = time.Now()
	return
}
func (m *Other) OnDelete() (status int, err error) {
	m.DeletedAt = time.Now()
	return
}
