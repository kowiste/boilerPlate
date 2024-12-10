package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Asset struct {
	id          string
	orgID       string
	name        string
	description string
	updatedAt   time.Time
	deletedAt   *time.Time
}

func New(orgID, name, description string) (*Asset, error) {
	if orgID == "" {
		return nil, ErrInvalidOrgID
	}
	if name == "" {
		return nil, ErrInvalidName
	}

	now := time.Now()
	return &Asset{
		id:          uuid.New().String(),
		orgID:       orgID,
		name:        name,
		description: description,
		updatedAt:   now,
	}, nil
}

func NewFromRepository(id, orgID, name, description string, updatedAt time.Time, deletedAt *time.Time) *Asset {
	return &Asset{
		id:          id,
		orgID:       orgID,
		name:        name,
		description: description,
		updatedAt:   updatedAt,
		deletedAt:   deletedAt,
	}
}

func (a *Asset) Update(name, description string, properties map[string]interface{}) error {
	if name == "" {
		return ErrInvalidName
	}

	a.name = name
	a.description = description
	a.updatedAt = time.Now()
	return nil
}

func (a *Asset) Delete() {
	now := time.Now()
	a.deletedAt = &now
}
func (a *Asset) IsDeleted() bool {
	return a.deletedAt != nil
}

// Getters
func (a *Asset) ID() string            { return a.id }
func (a *Asset) OrgID() string         { return a.orgID }
func (a *Asset) Name() string          { return a.name }
func (a *Asset) Description() string   { return a.description }
func (a *Asset) UpdatedAt() time.Time  { return a.updatedAt }
func (a *Asset) DeletedAt() *time.Time { return a.deletedAt }

type AssetRepository interface {
	Save(ctx context.Context, asset *Asset) error
	FindByID(ctx context.Context, orgID, assetID string) (*Asset, error)
	FindAll(ctx context.Context, orgID string) ([]*Asset, error)
	Remove(ctx context.Context, orgID, assetID string) error
}

var (
	ErrAssetNotFound = errors.New("asset not found")
	ErrInvalidOrgID  = errors.New("invalid organization id")
	ErrInvalidName   = errors.New("invalid name")
)
