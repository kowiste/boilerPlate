package infra

import (
	"context"
	"ddd/internal/features/asset/domain"
	"time"

	"gorm.io/gorm"
)

type assetRepository struct {
	db *gorm.DB
}

type Asset struct {
	ID          string `gorm:"primaryKey"`
	OrgID       string `gorm:"index"`
	Name        string
	Description string
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func NewRepository(db *gorm.DB) domain.AssetRepository {
	db.AutoMigrate(&Asset{})
	return &assetRepository{db: db}
}

func (r *assetRepository) Save(ctx context.Context, asset *domain.Asset) error {
	dbAsset := Asset{
		ID:          asset.ID(),
		OrgID:       asset.OrgID(),
		Name:        asset.Name(),
		Description: asset.Description(),
	}
	return r.db.WithContext(ctx).Save(&dbAsset).Error
}

func (r *assetRepository) FindByID(ctx context.Context, orgID, assetID string) (*domain.Asset, error) {
	var dbAsset Asset
	err := r.db.WithContext(ctx).Where("org_id = ? AND id = ?", orgID, assetID).First(&dbAsset).Error
	if err != nil {
		return nil, err
	}
	return domain.NewFromRepository(
		dbAsset.ID,
		dbAsset.OrgID,
		dbAsset.Name,
		dbAsset.Description,
		dbAsset.UpdatedAt,
		&dbAsset.DeletedAt,
	), nil

}

func (r *assetRepository) FindAll(ctx context.Context, orgID string) ([]*domain.Asset, error) {
	var dbAssets []Asset
	err := r.db.WithContext(ctx).Where("org_id = ?", orgID).Find(&dbAssets).Error
	if err != nil {
		return nil, err
	}

	assets := make([]*domain.Asset, len(dbAssets))
	for i, dbAsset := range dbAssets {
		assets[i] = domain.NewFromRepository(
			dbAsset.ID,
			dbAsset.OrgID,
			dbAsset.Name,
			dbAsset.Description,
			dbAsset.UpdatedAt,
			&dbAsset.DeletedAt,
		)
	}
	return assets, nil
}

func (r *assetRepository) Remove(ctx context.Context, orgID, assetID string) error {
	return r.db.WithContext(ctx).Where("org_id = ? AND id = ?", orgID, assetID).Delete(&Asset{}).Error
}
