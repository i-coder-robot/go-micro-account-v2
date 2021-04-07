package service

import (
	"errors"
	"github.com/i-coder-robot/go-micro-account-v2/model"
	"github.com/i-coder-robot/go-micro-account-v2/repository"
	"golang.org/x/crypto/bcrypt"
)

type AccountInterfaceService interface {
	AddAccount(account *model.Account) (int64, error)
	DeleteAccount(int64) error
	UpdateAccount(account *model.Account) (err error)
	FindAccountByName(string) (*model.Account, error)
	CheckPassword(accountName, password string) (suc bool, err error)
}

type AccountService struct {
	AccountRepository repository.AccountInterfaceRepository
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ValidatePassword(password string, hashed string) (isOK bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return false, errors.New("密码比较错误")
	}
	return true, nil
}

func (u AccountService) AddAccount(account *model.Account) (int64, error) {
	hashPwdByte, err := HashPassword(account.Md5Password)
	if err != nil {
		return account.ID, nil
	}
	account.Md5Password = string(hashPwdByte)
	return u.AccountRepository.CreateAccount(account)
}

func (u AccountService) DeleteAccount(id int64) error {
	return u.AccountRepository.DeleteAccount(id)
}

func (u AccountService) UpdateAccount(account *model.Account) (err error) {
	return u.AccountRepository.UpdateAccount(account)
}

func (u AccountService) FindAccountByName(name string) (*model.Account, error) {
	return u.AccountRepository.FindAccountByName(name)
}

func (u AccountService) CheckPassword(name, password string) (suc bool, err error) {
	account, err := u.AccountRepository.FindAccountByName(name)
	if err != nil {
		return false, err
	}
	return ValidatePassword(password, account.Md5Password)
}

func NewAccountService(repo repository.AccountInterfaceRepository) AccountInterfaceService {
	return &AccountService{
		AccountRepository: repo,
	}
}
