package middleware

import (
	"application_blog/common"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
)

func ThrowPanic() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				debugStack := ""

				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					debugStack += v + "\n"
				}

				// logger记录DebugStack,对接处理机制
				common.Log.ErrorMsg("[%s-%s-%s] %v :%s",
					c.Request.Method,
					c.Request.Host,
					c.Request.RequestURI,
					err,
					debugStack,
				)

				c.JSON(http.StatusInternalServerError, "系统错误")
				return
			}
		}()
		c.Next()
	}
}
