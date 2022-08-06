package services

import (
	"errors"
	"leafwearz/models/dto"
	"leafwearz/models/entities"
	"leafwearz/pkg/encryption"

	"gorm.io/gorm"
)

func (repo *AuthServicesResources) CreateUser(signupRequest dto.SignupRequest) (entities.User, error) {
	_, err := repo.FetchingRepos.FindUserByEmail(signupRequest.Email) 
	// If gorm does not find any existing users with specified email address, proceed
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newUser := entities.User{}
		newUser.Username = signupRequest.Username
		newUser.Email = signupRequest.Email
		newUser.Password = signupRequest.Password
		newUser, saveErr := repo.InsertionRepos.SaveUser(newUser)
		return newUser, saveErr
	} else {
		user := entities.User{}
		return user, errors.New("Problem occurred at data accessing layer: (user already exists!!!) failed to create user...")
	}
}

func (repo *AuthServicesResources) SigninUser(email string, password string) (uid uint, err error) {
	// check if user exists
	user, err := repo.FetchingRepos.FindUserByEmail(email)
	if err != nil {
		return 0, err
	}
	// check whether the password is correct
	if err := encryption.DecryptPassword(user.Password, password); err != nil {
		return 0, err
	}
	return user.ID, nil
}

