package usermodule

import (
	"letschat/common"
	"letschat/constant"
	"letschat/proto"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {

}

func registerUser(ctx *gin.Context) {
	req := new(proto.RegisterUserReq)
	if err := ctx.ShouldBindJSON(req); err != nil {
		common.LogUtil.Errorf("register user bind req error: %v", err)
		common.WriteRespone(ctx, nil, constant.Error_Requst)
		return
	}
	err := UserSvc.RegisterUser(req)
	common.WriteRespone(ctx, nil, err)
}

func (u *UserHandler) AddTo(e *gin.Engine) {
	group := e.Group("/user")
	{
		group.POST("/registerUser", registerUser)
	}
}