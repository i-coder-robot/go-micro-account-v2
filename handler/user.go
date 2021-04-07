package handler

import (
	"context"
	"github.com/i-coder-robot/go-micro-account-v2/model"
	account "github.com/i-coder-robot/go-micro-account-v2/proto"
	"github.com/i-coder-robot/go-micro-account-v2/service"
)

type User struct {
	AccountService service.AccountInterfaceService
}

func (u *User) Register(ctx context.Context, userRequest *account.AccountRequest, response *account.AccountResponse) error {
	userRegister := &model.Account{
		AccountName:    userRequest.AccountName,
		FirstName:   userRequest.FirstName,
		Md5Password: userRequest.Pwd,
	}
	_, err := u.AccountService.AddAccount(userRegister)
	if err != nil {
		return err
	}
	response.Message = "添加成功"
	return nil
}

func (u *User) GetUserInfo(ctx context.Context, userRequest *account.AccountRequest, response *account.AccountResponse) error {
	ok, err := u.AccountService.CheckPassword(userRequest.AccountName, userRequest.Pwd)
	if err != nil {
		return err
	}
	response.IsOk = ok
	return nil
}

func (u *User) Login(ctx context.Context, userRequest *account.AccountRequest, response *account.AccountResponse) error {
	info, err := u.AccountService.FindAccountByName(userRequest.AccountName)
	if err != nil {
		return err
	}

	response = convertUser4Response(info)

	return nil
}

func convertUser4Response(m *model.Account) *account.AccountResponse {
	res := &account.AccountResponse{}
	res.AccountName = m.AccountName
	res.FirstName = m.FirstName
	res.UserId = m.ID
	return res
}
