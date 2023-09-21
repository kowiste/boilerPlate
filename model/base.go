package model

import "time"

type BaseModel struct {
	ID        uint      `json:"id" binding:"omitempty,disabled" swaggerignore:"true"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `gorm:"index" json:"-"`
}

type ModelInterface interface {
	GetID() uint
	SetID(id uint)
	CreateValidation() (bool, map[string]string)
	UpdateValidation() (bool, map[string]string)
	BeforeValidation()
	AfterValidation()
}

func (m *BaseModel) CreateValidation() (bool, map[string]string) {
	return true, map[string]string{}
}

func (m *BaseModel) UpdateValidation() (bool, map[string]string) {
	return true, map[string]string{}
}

func (m *BaseModel) SetID(id uint) {
	m.ID = id
}

func (m *BaseModel) GetID() uint {
	return m.ID
}

func (m *BaseModel) BeforeValidation() {

}

func (m *BaseModel) AfterValidation() {

}
