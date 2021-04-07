package handler

import (
	"context"
	"github.com/i-coder-robot/go-micro-account-v2/model"
	account "github.com/i-coder-robot/go-micro-account-v2/proto"
	"github.com/i-coder-robot/go-micro-account-v2/service"
)

type Account struct {
	AccountService service.AccountInterfaceService
}

func (u *Account) Register(ctx context.Context, request *account.AccountRequest, response *account.AccountResponse) error {
	register := &model.Account{
		AccountName: request.AccountName,
		FirstName:   request.FirstName,
		Md5Password: request.Pwd,
	}
	_, err := u.AccountService.AddAccount(register)
	if err != nil {
		return err
	}
	response.Message = "添加成功"
	return nil
}

func (u *Account) GetAccountInfo(ctx context.Context, request *account.AccountRequest, response *account.AccountResponse) error {
	ok, err := u.AccountService.CheckPassword(request.AccountName, request.Pwd)
	if err != nil {
		return err
	}
	response.IsOk = ok
	return nil
}

func (u *Account) Login(ctx context.Context, request *account.AccountRequest, response *account.AccountResponse) error {
	info, err := u.AccountService.FindAccountByName(request.AccountName)
	if err != nil {
		return err
	}

	response = convertAccount4Response(info)

	return nil
}

func convertAccount4Response(m *model.Account) *account.AccountResponse {
	res := &account.AccountResponse{}
	res.AccountName = m.AccountName
	res.FirstName = m.FirstName
	res.AccountId = m.ID
	return res
}
