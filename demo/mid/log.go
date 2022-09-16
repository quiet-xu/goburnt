package mid

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Log(c *gin.Context) {
	fmt.Println("LOG MID")
}

func Auth(c *gin.Context) {
	fmt.Println("Auth MID")
}
