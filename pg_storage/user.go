package pg_storage

import (
	"github.com/ilhamrobyana/bts-test/entity"
	"github.com/jinzhu/gorm"
)

type User struct{}

func (u *User) Create(user entity.User) (entity.User, error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return entity.User{}, e
	}

	registered := checkRegisteredByEmail(client, user.Email)
	if !registered {
		e = client.Create(&user).Error
		if e != nil {
			return entity.User{}, e
		}
	}
	return user, nil
}

func (u *User) GetByEmail(email string) (user entity.User, e error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return
	}
	e = client.Where("email = ?", email).First(&user).Error
	return
}

func (u *User) GetAll() (users []entity.User, e error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return
	}
	e = client.
		Find(&users).
		Order("id ASC").
		Error
	return
}

func checkRegisteredByEmail(client *gorm.DB, email string) bool {
	user := entity.User{}
	e := client.Where("email=?", email).Find(&user)
	if e == nil {
		return true
	}
	return false
}
