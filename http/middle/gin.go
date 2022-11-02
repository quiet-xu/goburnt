package middle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	gin2 "github.com/quiet-xu/goburnt/http/gin"
)

type GinMiddle struct {
	authKey         string
	cbt             *gin2.Cbt
	authHandlerData any
}

func NewGinMiddle() *GinMiddle {
	return &GinMiddle{
		authKey: "Authorization",
		cbt:     gin2.NewCbt(),
	}
}

// SetAuthKey 设置auth key
func (s *GinMiddle) SetAuthKey(authKey string) *GinMiddle {
	s.authKey = authKey
	return s
}

func (s *GinMiddle) Auth(ctx *gin.Context) {
	token := ctx.Request.Header.Get(s.authKey)
	if len(token) == 0 {
		s.cbt.FailWithData(fmt.Errorf("鉴权信息不能为空"), ctx)
		ctx.Abort()
		return
	}
	return
}
