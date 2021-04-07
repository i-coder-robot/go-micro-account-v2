package model

type Account struct {
	ID          int64  `gorm:"primary_key;not_null;auto_increment"`
	AccountName string `gorm："unique_index;not_null"`
	FirstName   string
	Md5Password string
}
