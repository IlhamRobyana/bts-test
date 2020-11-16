package pg_storage

import "github.com/ilhamrobyana/bts-test/entity"

type Shopping struct{}

func (s *Shopping) Create(shopping entity.Shopping) (entity.Shopping, error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return shopping, e
	}

	e = client.Create(&shopping).Error
	return shopping, e
}
func (s *Shopping) GetAll() (shoppingList []entity.Shopping, e error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return
	}

	e = client.
		Find(&shoppingList).
		Order("id ASC").
		Error
	return
}
func (s *Shopping) GetByID(id uint64) (shopping entity.Shopping, e error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return
	}

	e = client.
		Where("id=? ", id).
		Find(&shopping).
		Error
	return
}
func (s *Shopping) Update(id uint64, shopping entity.Shopping) (entity.Shopping, error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return shopping, e
	}

	e = client.
		Model(&shopping).
		Where("id=?", id).
		Updates(shopping).
		Where("id=?", id).
		Find(&shopping).Error
	return shopping, e
}
func (s *Shopping) Delete(id uint64) (e error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return
	}
	e = client.
		Where("id=?", id).
		Find(&entity.Shopping{}).
		Delete(&entity.Shopping{}).
		Error
	return
}
