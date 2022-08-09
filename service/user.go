package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/status"
	"gin_mall/pkg/utils"
	"gin_mall/serializer"
)

// UserService 用户服务
type UserService struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
}

func (service UserService) Register(ctx context.Context) serializer.Response { // serializer.Response是自定义的返回给前端的响应体结构
	code := status.SUCCESS
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = status.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    status.GetMsg(code),
		}
	}
	if exist {
		code = status.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    status.GetMsg(code),
		}
	}
	user := &model.User{
		NickName: service.NickName,
		UserName: service.UserName,
		Status:   model.Active,
		Money:    utils.Encrypt.AesEncoding("10000"), // 初始金额
	}
	//加密密码
	if err = user.SetPassword(service.Password); err != nil {
		code = status.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    status.GetMsg(code),
		}
	}
	user.Avatar = "http://q1.qlogo.cn/g?b=qq&nk=294350394&s=640"
	//创建用户
	err = userDao.CreateUser(user)
	if err != nil {
		code = status.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    status.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    status.GetMsg(code),
	}
}

//Login 用户登陆函数
func (service UserService) Login(ctx context.Context) serializer.Response {
	code := status.SUCCESS
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if !exist { //如果查询不到，返回相应的错误
		code = status.ErrorUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    status.GetMsg(code),
		}
	}
	if user.CheckPassword(service.Password) == false {
		code = status.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    status.GetMsg(code),
		}
	}
	// token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		code = status.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    status.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		// Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg: status.GetMsg(code),
	}
}
