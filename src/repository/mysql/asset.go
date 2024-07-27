package mysql

import (
	"context"

	"github.com/kowiste/boilerplate/src/model/asset"
)

func (m MySQL) CreateAsset(c context.Context, asset *asset.Asset) (id string, err error) {
	result := m.db.Create(asset)
	if result.Error != nil {
		err = result.Error
		return
	}
	id = asset.ID
	return
}

func (m MySQL) Assets(c context.Context) (assets asset.Assets, err error) {
	assets = make([]asset.Asset, 0)
	result := m.db.Find(&assets)
	if result.Error != nil {
		return nil, result.Error
	}
	return
}
func (m MySQL) AssetByID(c context.Context, id string) (asset *asset.Asset, err error) {
	result := m.db.Where("id = ?", id).First(&asset)
	if result.Error != nil {
		return nil, result.Error
	}
	return
}
func (m MySQL) UpdateAsset(c context.Context, asset *asset.Asset) (err error) {
	result := m.db.Model(asset).Where("id = ?", asset.ID).Updates(asset)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
func (m MySQL) DeleteAsset(c context.Context, id string) (err error) {
	result := m.db.Delete(&asset.Asset{}, "id = ?", id)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
