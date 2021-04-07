package service

import (
	"errors"
	"github.com/i-coder-robot/go-micro-account-v2/model"
	"github.com/i-coder-robot/go-micro-account-v2/repository"
	"golang.org/x/crypto/bcrypt"
)

type AccountInterfaceService interface {
	AddAccount(user *model.Account) (int64, error)
	DeleteAccount(int64) error
	UpdateAccount(user *model.Account) (err error)
	FindAccountByName(string) (*model.Account, error)
	CheckPassword(accountName, password string) (suc bool, err error)
}

type UserService struct {
	AccountRepository repository.AccountInterfaceRepository
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

func (u UserService) AddAccount(account *model.Account) (int64, error) {
	hashPwdByte, err := HashPassword(account.Md5Password)
	if err != nil {
		return account.ID, nil
	}
	account.Md5Password = string(hashPwdByte)
	return u.AccountRepository.CreateAccount(account)
}

func (u UserService) DeleteAccount(id int64) error {
	return u.AccountRepository.DeleteAccount(id)
}

func (u UserService) UpdateAccount(account *model.Account) (err error) {
	return u.AccountRepository.UpdateAccount(account)
}

func (u UserService) FindAccountByName(name string) (*model.Account, error) {
	return u.AccountRepository.FindAccountByName(name)
}

func (u UserService) CheckPassword(name, password string) (suc bool, err error) {
	user, err := u.AccountRepository.FindAccountByName(name)
	if err != nil {
		return false, err
	}
	return ValidatePassword(password, user.Md5Password)
}

func NewUserService(repo repository.AccountInterfaceRepository) AccountInterfaceService {
	return &UserService{
		AccountRepository: repo,
	}
}
