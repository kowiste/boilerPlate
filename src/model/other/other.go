package other

import (
	"serviceX/src/model"
	"time"
)

type Other struct {
	model.BaseNoSQL `bson:"inline"`
	controller      model.ControllerI `json:"-"`
	Field1          int               `json:"field1" example:"1" binding:"max=200"`
	Name            string            `json:"name" example:"peter"`
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
func (m *Other) OnCreate() {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

}
func (m *Other) OnUpdate() {
	m.UpdatedAt = time.Now()
}
func (m *Other) OnDelete() {
	m.DeletedAt = time.Now()
}
