package shopping

import (
	"github.com/ilhamrobyana/bts-test/entity"
	"github.com/ilhamrobyana/bts-test/storage"
)

type core struct {
	shoppingStorage storage.ShoppingStorage
}

func (c *core) create(shopping entity.Shopping) (createdShopping entity.Shopping, e error) {
	createdShopping, e = c.shoppingStorage.Create(shopping)
	return
}

func (c *core) getAll() (shoppingList []entity.Shopping, e error) {
	shoppingList, e = c.shoppingStorage.GetAll()
	return
}

func (c *core) getByID(id uint64) (shopping entity.Shopping, e error) {
	shopping, e = c.shoppingStorage.GetByID(id)
	return
}

func (c *core) update(id uint64, shopping entity.Shopping) (entity.Shopping, error) {
	shopping, e := c.shoppingStorage.Update(id, shopping)
	return shopping, e
}

func (c *core) delete(id uint64) (e error) {
	e = c.shoppingStorage.Delete(id)
	return
}
