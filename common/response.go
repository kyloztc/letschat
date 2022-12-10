package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func WriteRespone(ctx *gin.Context, data interface{}, err error) {
	response := &BaseResponse{
		Code: 0,
		Message: "success",
		Data: data,
	}
	if err != nil {
		response.Code = -1
		response.Message = err.Error()
	}
	ctx.JSON(http.StatusOK, response)
}