package user

import (
	"application_blog/common"
	"application_blog/conf"
	"application_blog/internal/pkg/crypto"
	client "application_blog/internal/pkg/grpc"
	pb "application_blog/internal/protocal/pb/blog"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type ChangePasswordParam struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func ChangePassword(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	param := ChangePasswordParam{}
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErrInvalidParam("BindJSON err:"+err.Error()))
		return
	}
	ctx, cancel := common.SetContextWithTimeout(c, conf.GrpcTimeOut)
	defer cancel()
	oldPassword, err := crypto.DecryptApiData(param.OldPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErr("旧密码解密错误"))
		return
	}
	newPassword, err := crypto.DecryptApiData(param.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErr("新密码解密错误"))
		return
	}
	_, err = client.GrpcClients.BlogHostClient.ChangePassword(ctx, &pb.ChangePasswordRequest{
		Username:    param.Username,
		OldPassword: oldPassword,
		NewPassword: newPassword,
		Token:       token,
	})
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "不存在该用户") {
			c.JSON(http.StatusInternalServerError, common.NewErr("不存在该用户"))
		} else if strings.Contains(errMsg, "旧密码错误") {
			c.JSON(http.StatusInternalServerError, common.NewErr("旧密码错误"))
		} else if strings.Contains(errMsg, "新旧密码不能一样") {
			c.JSON(http.StatusInternalServerError, common.NewErr("新旧密码不能一样"))
		} else {
			c.JSON(http.StatusInternalServerError, common.NewErr(errMsg))
		}
		return
	}
	c.JSON(http.StatusOK, common.ResponseSuccess(nil))
	return
}
