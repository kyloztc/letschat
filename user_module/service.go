package usermodule

import (
	"letschat/constant"
	"letschat/dao"
	"letschat/proto"
)

var UserSvc = new(UserService)

type UserService struct {

}

func (s *UserService) RegisterUser(req *proto.RegisterUserReq) error {
	user, err := dao.UserDaoUtil.QueryUserByAccount(req.Account)
	if err != nil && err != constant.Error_Data_Not_Found {
		return err
	}
	if user != nil {
		return constant.Error_User_Registed
	}
	user = &dao.User{
		Account: req.Account,
		Pwd: req.Pwd,
		NickName: req.NickName,
	}
	
	_, err = dao.UserDaoUtil.CreateUser(user)
	return err
}