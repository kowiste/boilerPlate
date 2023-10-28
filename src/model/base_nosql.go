package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseNoSQL struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty" binding:"omitempty" swaggerignore:"true"`
	CreatedAt time.Time          `json:"-"`
	UpdatedAt time.Time          `json:"-"`
	DeletedAt time.Time          `json:"-" bson:"omitempty"`
}

func (m *BaseNoSQL) CreateValidation() (bool, map[string]string) {
	return true, map[string]string{}
}

func (m *BaseNoSQL) UpdateValidation() (bool, map[string]string) {
	return true, map[string]string{}
}

func (m *BaseNoSQL) SetID(id string) (err error) {
	m.ID, err = primitive.ObjectIDFromHex(id)
	return
}

func (m *BaseNoSQL) GetID() string {
	return m.ID.String()
}

// BeforeValidation remove id from validation
func (m *BaseNoSQL) BeforeValidation() {

}

// AfterValidation add id after validation
func (m *BaseNoSQL) AfterValidation() {

}
