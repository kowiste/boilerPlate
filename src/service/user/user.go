package userservice

import (
	"sync"

	"github.com/kowiste/boilerplate/pkg/errors"
	"github.com/kowiste/boilerplate/src/db"
	"github.com/kowiste/boilerplate/src/messaging"
	"github.com/kowiste/boilerplate/src/model/user"
	"github.com/kowiste/boilerplate/src/transport"
)

type UserService struct {
	user      *user.User
	db        db.IDatabase
	msg       messaging.IMessaging
	transport transport.ITransport
}
type Option func(*UserService)

var (
	instance *UserService
	once     sync.Once
)

func New(opts ...Option) (serv *UserService) {

	once.Do(func() {

		instance = &UserService{
			user: new(user.User),
		}
		for _, opt := range opts {
			opt(instance)
		}
	})
	return instance
}
func WithDatabase(db db.IDatabase) Option {
	return func(s *UserService) {
		s.db = db
	}
}

func WithMessaging(msg messaging.IMessaging) Option {
	return func(s *UserService) {
		s.msg = msg
	}
}

func WithTransport(transport transport.ITransport) Option {
	return func(s *UserService) {
		s.transport = transport
	}
}
func Get() (*UserService, error) {
	if instance == nil {
		return nil, errors.New("AssetService not set", errors.EErrorServerInternal)
	}
	return instance, nil
}
func (serv *UserService) GetUser() *user.User {
	return serv.user
}
