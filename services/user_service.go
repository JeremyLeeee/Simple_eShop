package services

import (
	"eshop/datamodels"
	"eshop/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	IsPwdSuccess(userName string, pwd string) (user *datamodels.User, isOk bool)
	GetUserByName(string) (*datamodels.User, error)
	AddUser(*datamodels.User) (int64, error)
}

type UserService struct {
	userRepository repositories.IUser
}

func NewUserService(userRepository repositories.IUser) IUserService {
	return &UserService{userRepository}
}

func (s *UserService) IsPwdSuccess(userName string, pwd string) (user *datamodels.User, isOk bool) {
	user, err := s.userRepository.Select(userName)
	if err != nil {
		return &datamodels.User{}, false
	}
	isOK, _ := validatePassword(pwd, user.HashPassword)
	if !isOK {
		return &datamodels.User{}, false
	}
	return
}

func (s *UserService) GetUserByName(userName string) (user *datamodels.User, err error) {
	return s.userRepository.Select(userName)
}

func (s *UserService) AddUser(user *datamodels.User) (ID int64, err error) {
	pwdByte, err := generatePassword(user.HashPassword)
	if err != nil {
		return ID, err
	}
	user.HashPassword = string(pwdByte)
	return s.userRepository.Insert(user)
}

func validatePassword(pwd string, hashedPwd string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
	if err != nil {
		return false, err
	}
	return true, nil
}

func generatePassword(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}
