package storage

import (
	"errors"

	"github.com/ilhamrobyana/bts-test/entity"
	pg "github.com/ilhamrobyana/bts-test/pg_storage"
)

type ShoppingStorage interface {
	Create(shopping entity.Shopping) (entity.Shopping, error)
	GetAll() ([]entity.Shopping, error)
	GetByID(id uint64) (entity.Shopping, error)
	Update(id uint64, shopping entity.Shopping) (entity.Shopping, error)
	Delete(id uint64) error
}

func GetShoppingStorage(n int) (ShoppingStorage, error) {
	switch n {
	case Postgre:
		return new(pg.Shopping), nil
	default:
		return nil, errors.New("not implemented yet")
	}
}
