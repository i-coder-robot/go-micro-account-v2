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
	UpdateAccount(user *model.Account) error
	FindAll() ([]model.Account, error)
}

func NewUserInterfaceRepository(db *gorm.DB) AccountInterfaceRepository {
	return &UserRepository{
		db: db,
	}
}

type UserRepository struct {
	db *gorm.DB
}

func (u UserRepository) InitTable() error {
	exist := u.db.HasTable(&model.Account{})
	if exist {
		return nil
	}
	return u.db.CreateTable(&model.Account{}).Error
}

func (u UserRepository) FindAccountByName(name string) (*model.Account, error) {
	user := &model.Account{}
	err := u.db.Where("user_name = ?", name).Find(user).Error
	return user, err
}

func (u UserRepository) FindAccountById(id int64) (*model.Account, error) {
	user := &model.Account{}
	return user, u.db.First(user, id).Error
}

func (u UserRepository) CreateAccount(user *model.Account) (int64, error) {
	return user.ID, u.db.Create(user).Error
}

func (u UserRepository) DeleteAccount(id int64) error {
	return u.db.Where("id = ?", id).Delete(&model.Account{}).Error
}

func (u UserRepository) UpdateAccount(user *model.Account) error {
	return u.db.Model(user).Update(user).Error
}

func (u UserRepository) FindAll() (users []model.Account, err error) {
	return users, u.db.Find(&users).Error
}
