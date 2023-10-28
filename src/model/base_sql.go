package model

import (
	"time"
)

type BaseSQL struct {
	ID        uint      `json:"id" gorm:"id"  binding:"omitempty" swaggerignore:"true"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `gorm:"index" json:"-"`
	id        uint      //temp id storage
}

func (m *BaseSQL) CreateValidation() (bool, map[string]string) {
	return true, map[string]string{}
}

func (m *BaseSQL) UpdateValidation() (bool, map[string]string) {
	return true, map[string]string{}
}

func (m *BaseSQL) setID(id uint) {
	m.ID = id
}

func (m *BaseSQL) GetID() string {
	return string(m.ID)
}
func (m *BaseSQL) getID() uint {
	return m.ID
}

// BeforeValidation remove id from validation
func (m *BaseSQL) BeforeValidation() {
	m.id = m.getID()
	m.setID(0)
}

// AfterValidation add id after validation
func (m *BaseSQL) AfterValidation() {
	m.setID(m.id)
}
