package service

import (
	"testApi/model"
	"testApi/repository"
)

func Authentificate(username string, password string) (model.User, bool, error) {

	user, ok, err := repository.Login(username, password)

	return user, ok, err
}