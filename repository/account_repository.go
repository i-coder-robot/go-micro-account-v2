package repository

import (
	"github.com/i-coder-robot/go-micro-account-v2/model"
	"github.com/jinzhu/gorm"
)

type AccountInterfaceRepository interface {
	InitTable() error
	FindAccountByName(string) (*model.Account, error)
	FindAccountById(int642 int64) (*model.Account, error)
	CreateAccount(*model.Account) (int64, error)
	DeleteAccount(int64) error
	UpdateAccount(account *model.Account) error
	FindAll() ([]model.Account, error)
}

func NewAccountInterfaceRepository(db *gorm.DB) AccountInterfaceRepository {
	return &AccountRepository{
		db: db,
	}
}

type AccountRepository struct {
	db *gorm.DB
}

func (u AccountRepository) InitTable() error {
	exist := u.db.HasTable(&model.Account{})
	if exist {
		return nil
	}
	return u.db.CreateTable(&model.Account{}).Error
}

func (u AccountRepository) FindAccountByName(name string) (*model.Account, error) {
	account := &model.Account{}
	err := u.db.Where("account_name = ?", name).Find(account).Error
	return account, err
}

func (u AccountRepository) FindAccountById(id int64) (*model.Account, error) {
	account := &model.Account{}
	return account, u.db.First(account, id).Error
}

func (u AccountRepository) CreateAccount(account *model.Account) (int64, error) {
	return account.ID, u.db.Create(account).Error
}

func (u AccountRepository) DeleteAccount(id int64) error {
	return u.db.Where("id = ?", id).Delete(&model.Account{}).Error
}

func (u AccountRepository) UpdateAccount(account *model.Account) error {
	return u.db.Model(account).Update(account).Error
}

func (u AccountRepository) FindAll() (accounts []model.Account, err error) {
	return accounts, u.db.Find(&accounts).Error
}
