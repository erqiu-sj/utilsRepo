package gin

import "github.com/gin-gonic/gin"

type Decorator struct {
}

func (receiver Decorator) PermissionCharacters(permission string, fn func(context *gin.Context)) func(context *gin.Context) {
	return fn
}
