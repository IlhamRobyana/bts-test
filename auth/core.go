package auth

import (
	"errors"

	"github.com/ilhamrobyana/bts-test/entity"
	"github.com/ilhamrobyana/bts-test/helper"
	"github.com/ilhamrobyana/bts-test/storage"
	"golang.org/x/crypto/bcrypt"
)

type core struct {
	userStorage  storage.UserStorage
	tokenStorage storage.TokenStorage
}

func (c *core) signup(user entity.User) (response entity.LoginResponse, e error) {
	hashedPassword, e := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashedPassword)

	createdUser, e := c.userStorage.Create(user)
	token, e := helper.GenerateToken(user)
	if e != nil {
		return
	}
	e = c.tokenStorage.Create(token)
	response = entity.LoginResponse{}
	response.Email = createdUser.Email
	response.Token = token
	response.Username = createdUser.Username
	return
}

func (c *core) signin(email, password string) (response entity.LoginResponse, e error) {
	user, e := c.userStorage.GetByEmail(email)
	if e != nil {
		return
	}
	e = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if e != nil {
		e = errors.New("Email or Password is wrong")
	}

	token, e := helper.GenerateToken(user)
	if e != nil {
		return
	}

	e = c.tokenStorage.Create(token)
	response = entity.LoginResponse{}
	response.Email = user.Email
	response.Token = token
	response.Username = user.Username
	return
}

func (c *core) getAll() (users []entity.User, e error) {
	users, e = c.userStorage.GetAll()
	return
}
