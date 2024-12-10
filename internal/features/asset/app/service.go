package app

import (
	"context"
	"ddd/internal/features/asset/domain"
	"ddd/shared/logger"
	"fmt"
)

type AssetService interface {
	CreateAsset(ctx context.Context, cmd CreateAssetCommand) (*domain.Asset, error)
	GetAsset(ctx context.Context, orgID, assetID string) (*domain.Asset, error)
	ListAssets(ctx context.Context, orgID string) ([]*domain.Asset, error)
	UpdateAsset(ctx context.Context, cmd UpdateAssetCommand) (*domain.Asset, error)
	DeleteAsset(ctx context.Context, orgID, assetID string) error
}
type assetService struct {
	repo   domain.AssetRepository
	logger logger.Logger
}

func NewService(repo domain.AssetRepository, logger logger.Logger) AssetService {
	return &assetService{
		repo:   repo,
		logger: logger,
	}
}
func (s *assetService) CreateAsset(ctx context.Context, cmd CreateAssetCommand) (*domain.Asset, error) {
	asset, err := domain.New(cmd.OrgID, cmd.Name, cmd.Description)
	if err != nil {
		return nil, fmt.Errorf("failed to create asset: %w", err)
	}
	if err := s.repo.Save(ctx, asset); err != nil {
		return nil, fmt.Errorf("failed to create asset: %w", err)
	}

	return asset, nil
}

func (s *assetService) GetAsset(ctx context.Context, orgID, assetID string) (*domain.Asset, error) {
	asset, err := s.repo.FindByID(ctx, orgID, assetID)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %w", err)
	}
	return asset, nil
}

func (s *assetService) ListAssets(ctx context.Context, orgID string) ([]*domain.Asset, error) {
	assets, err := s.repo.FindAll(ctx, orgID)
	if err != nil {
		return nil, fmt.Errorf("failed to list assets: %w", err)
	}
	return assets, nil
}

func (s *assetService) UpdateAsset(ctx context.Context, cmd UpdateAssetCommand) (*domain.Asset, error) {
	asset, err := s.repo.FindByID(ctx, cmd.OrgID, cmd.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %w", err)
	}
	err = asset.Update(cmd.Name, cmd.Description, cmd.Properties)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %w", err)
	}
	if err := s.repo.Save(ctx, asset); err != nil {
		return nil, fmt.Errorf("failed to update asset: %w", err)
	}

	return asset, nil
}
func (s *assetService) DeleteAsset(ctx context.Context, orgID, assetID string) error {
	asset, err := s.repo.FindByID(ctx, orgID, assetID)
	if err != nil {
		return fmt.Errorf("failed to find asset: %w", err)
	}

	asset.Delete()
	if err := s.repo.Save(ctx, asset); err != nil {
		return fmt.Errorf("failed to save deleted asset: %w", err)
	}
	return nil
}
