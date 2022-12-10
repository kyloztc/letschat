package dao

import (
	"fmt"
	"letschat/common"
	"testing"
)

var configPath = "../conf/properties.yaml"

func TestCreatUser(t *testing.T) {
	common.Init(configPath)
	user := &User{
		Account:  "testAccount",
		Pwd:      "pwd",
		NickName: "testAccount",
	}
	id, err := UserDaoUtil.CreateUser(user)
	fmt.Printf("id: %v|err: %v\n", id, err)
}
