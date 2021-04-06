package repository

import (
	"github.com/i-coder-robot/go-micro-user-v2/model"
	"github.com/jinzhu/gorm"
)

type UserInterfaceRepository interface {
	InitTable() error
	FindUserByName(string) (*model.User, error)
	FindUserById(int642 int64) (*model.User, error)
	CreateUser(*model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(user *model.User) error
	FindAll() ([]model.User, error)
}

func NewUserInterfaceRepository(db *gorm.DB) UserInterfaceRepository {
	return &UserRepository{
		db: db,
	}
}

type UserRepository struct {
	db *gorm.DB
}

func (u UserRepository) InitTable() error {
	exist := u.db.HasTable(&model.User{})
	if exist {
		return nil
	}
	return u.db.CreateTable(&model.User{}).Error
}

func (u UserRepository) FindUserByName(name string) (*model.User, error) {
	user := &model.User{}
	err := u.db.Where("user_name = ?", name).Find(user).Error
	return user, err
}

func (u UserRepository) FindUserById(id int64) (*model.User, error) {
	user := &model.User{}
	return user, u.db.First(user, id).Error
}

func (u UserRepository) CreateUser(user *model.User) (int64, error) {
	return user.ID, u.db.Create(user).Error
}

func (u UserRepository) DeleteUser(id int64) error {
	return u.db.Where("id = ?", id).Delete(&model.User{}).Error
}

func (u UserRepository) UpdateUser(user *model.User) error {
	return u.db.Model(user).Update(user).Error
}

func (u UserRepository) FindAll() (users []model.User, err error) {
	return users, u.db.Find(&users).Error
}
