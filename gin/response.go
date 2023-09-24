package gin

import (
	"github.com/gin-gonic/gin"
	"memo28.com/repo/utils"
	"net/http"
)

type Response struct {
}

// Suc 处理成功返回值
func (receiver Response) Suc(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, receiver.EncapsulatesResponse(http.StatusOK, data, "ok"))
	context.Abort()
}

// Error 处理错误返回
func (receiver Response) Error(context *gin.Context, code int, data interface{}, msg string) {
	context.JSON(http.StatusOK, receiver.EncapsulatesResponse(code, data, msg))
	context.Abort()
}

// EncapsulatesResponse 组装响应数据
func (receiver Response) EncapsulatesResponse(code int, data interface{}, msg string) utils.Map[string, any] {
	u := utils.Map[string, any]{}
	u.Add("code", code).Add("data", data).Add("msg", msg)
	return u
}
