package assetservice

import (
	"sync"

	"github.com/kowiste/boilerplate/pkg/errors"
	"github.com/kowiste/boilerplate/src/db"
	"github.com/kowiste/boilerplate/src/messaging"
	"github.com/kowiste/boilerplate/src/model/asset"
	"github.com/kowiste/boilerplate/src/transport"
)

type AssetService struct {
	asset     *asset.Asset
	db        db.IDatabase
	msg       messaging.IMessaging
	transport transport.ITransport
}
type Option func(*AssetService)

var (
	instance *AssetService
	once     sync.Once
)

func New(opts ...Option) (serv *AssetService) {

	once.Do(func() {

		instance = &AssetService{
			asset: new(asset.Asset),
		}
		for _, opt := range opts {
			opt(instance)
		}
	})

	return instance
}
func WithDatabase(db db.IDatabase) Option {
	return func(s *AssetService) {
		s.db = db
	}
}

func WithMessaging(msg messaging.IMessaging) Option {
	return func(s *AssetService) {
		s.msg = msg
	}
}

func WithTransport(transport transport.ITransport) Option {
	return func(s *AssetService) {
		s.transport = transport
	}
}

func Get() (*AssetService, error) {
	if instance == nil {
		return nil, errors.New("AssetService not set", errors.EErrorServerInternal)
	}
	return instance, nil
}

func (serv *AssetService) GetAsset() *asset.Asset {
	return serv.asset
}
