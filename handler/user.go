package handler

import (
	"context"
	"github.com/i-coder-robot/go-micro-user-v2/model"
	user "github.com/i-coder-robot/go-micro-user-v2/proto"
	"github.com/i-coder-robot/go-micro-user-v2/service"
)

type User struct {
	UserService service.UserInterfaceService
}

func (u *User) Register(ctx context.Context, userRequest *user.UserRequest, response *user.UserResponse) error {
	userRegister := &model.User{
		UserName:    userRequest.UserName,
		FirstName:   userRequest.FirstName,
		Md5Password: userRequest.Pwd,
	}
	_, err := u.UserService.AddUser(userRegister)
	if err != nil {
		return err
	}
	response.Message = "添加成功"
	return nil
}

func (u *User) GetUserInfo(ctx context.Context, userRequest *user.UserRequest, response *user.UserResponse) error {
	ok, err := u.UserService.CheckPassword(userRequest.UserName, userRequest.Pwd)
	if err != nil {
		return err
	}
	response.IsOk = ok
	return nil
}

func (u *User) Login(ctx context.Context, userRequest *user.UserRequest, response *user.UserResponse) error {
	info, err := u.UserService.FindUserByName(userRequest.UserName)
	if err != nil {
		return err
	}

	response = convertUser4Response(info)

	return nil
}

func convertUser4Response(userModel *model.User) *user.UserResponse {
	res := &user.UserResponse{}
	res.UserName = userModel.UserName
	res.FirstName = userModel.FirstName
	res.UserId = userModel.ID
	return res
}
