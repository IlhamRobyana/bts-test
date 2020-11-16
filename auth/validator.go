package auth

import (
	"errors"

	"github.com/ilhamrobyana/bts-test/entity"
)

func checkLogin(request entity.LoginRequest) error {
	if len(request.Email) == 0 {
		return errors.New("Phone must not be empty")
	}
	if len(request.Password) == 0 {
		return errors.New("Password must not be empty")
	}
	return nil
}
