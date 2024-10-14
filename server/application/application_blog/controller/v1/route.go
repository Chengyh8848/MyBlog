package v1

import (
	"application_blog/controller/v1/user"

	"github.com/gin-gonic/gin"
)

func NoNeedLoginRoute(g *gin.RouterGroup) {
	// 登录
	g.POST("/user/login", user.LoginUser)

}

func OnlyNeedLoginRoute(g *gin.RouterGroup) {
	g = g.Group("/v1")

	g.POST("/user/logout", user.Logout)
	g.POST("/user/changePassword", user.ChangePassword)
}
