package service

import (
	"errors"
	"github.com/i-coder-robot/go-micro-user-v2/model"
	"github.com/i-coder-robot/go-micro-user-v2/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserInterfaceService interface {
	AddUser(user *model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(user *model.User) (err error)
	FindUserByName(string) (*model.User, error)
	CheckPassword(userName, password string) (suc bool, err error)
}

type UserService struct {
	UserRepository repository.UserInterfaceRepository
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string, hashed string) (isOK bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword))
	if err != nil {
		return false, errors.New("密码比较错误")
	}
	return true, nil
}

func (u UserService) AddUser(user *model.User) (int64, error) {
	hashPwdByte, err := HashPassword(user.Md5Password)
	if err != nil {
		return user.ID, nil
	}
	user.Md5Password = string(hashPwdByte)
	return u.UserRepository.CreateUser(user)
}

func (u UserService) DeleteUser(userId int64) error {
	return u.UserRepository.DeleteUser(userId)
}

func (u UserService) UpdateUser(user *model.User) (err error) {
	return u.UserRepository.UpdateUser(user)
}

func (u UserService) FindUserByName(userName string) (*model.User, error) {
	return u.UserRepository.FindUserByName(userName)
}

func (u UserService) CheckPassword(userName, password string) (suc bool, err error) {
	user, err := u.UserRepository.FindUserByName(userName)
	if err != nil {
		return false, err
	}
	return ValidatePassword(password, user.Md5Password)
}

func NewUserService(repo repository.UserInterfaceRepository) UserInterfaceService {
	return &UserService{
		UserRepository: repo,
	}
}
