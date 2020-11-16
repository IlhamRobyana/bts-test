package storage

import (
	"errors"

	"github.com/ilhamrobyana/bts-test/entity"
	pg "github.com/ilhamrobyana/bts-test/pg_storage"
)

type UserStorage interface {
	Create(user entity.User) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
	GetAll() ([]entity.User, error)
}

func GetUserStorage(n int) (UserStorage, error) {
	switch n {
	case Postgre:
		return new(pg.User), nil
	default:
		return nil, errors.New("not implemented yet")
	}
}
