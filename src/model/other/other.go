package other

import (
	"serviceX/src/model"
)

type Stuff struct {
	model.BaseSQL
	controller model.ControllerI `json:"-"`
	Field1     int               `json:"field1" example:"1"`
	Name       string            `json:"name" example:"peter"`
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
