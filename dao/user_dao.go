package dao

import (
	"letschat/common"
	"letschat/constant"
	"letschat/util/encrypt"
)

type User struct {
	Id       int    `gorm:"primaryKey;column:f_id"`
	Account  string `gorm:"column:f_account"`
	Pwd      string `gorm:"column:f_pwd"`
	NickName string `gorm:"column:f_nick_name"`
}

var UserDaoUtil = new(UserDao)

type UserDao struct {
}

func (d *UserDao) CreateUser(user *User) (int, error) {
	if user == nil {
		return 0, constant.Error_Unknow
	}
	encodePwd := encrypt.Md5Encrypt(user.Pwd)
	user.Pwd = encodePwd
	result := common.DbConn.Create(user)
	id, err := user.Id, result.Error
	if err != nil {
		common.LogUtil.Errorf("create data error: %v", err)
		return -1, err
	}
	return id, nil
}

func (d *UserDao) QueryUserByAccount(account string) (*User, error) {
	user := new(User)
	whereMap := map[string]interface{}{
		"f_account": account,
	}
	common.DbConn.Where(whereMap).First(user)

	if user.Id == 0 {
		common.LogUtil.Errorf("data not found|account: %v", account)
		return nil, constant.Error_Data_Not_Found
	}
	return user, nil
}
