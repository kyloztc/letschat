package constant

import "errors"

var (
	Error_Data_Not_Found = errors.New("data not found")
	Error_User_Registed = errors.New("account has been registered")
	Error_Requst = errors.New("parse request error")
	Error_Unknow = errors.New("unknow error")
)