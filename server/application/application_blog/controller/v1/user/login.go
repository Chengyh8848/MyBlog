package user

import (
	"application_blog/common"
	"application_blog/conf"
	"application_blog/internal/pkg/crypto"
	client "application_blog/internal/pkg/grpc"
	pb "application_blog/internal/protocal/pb/blog"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginUserParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUser(c *gin.Context) {
	param := LoginUserParam{}
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErrInvalidParam("BindJSON err:"+err.Error()))
		return
	}
	ctx, cancel := common.SetContextWithTimeout(c, conf.GrpcTimeOut)
	defer cancel()
	password, err := crypto.DecryptApiData(param.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErr("密码解密错误"))
		return
	}
	reply, err := client.GrpcClients.BlogHostClient.LoginUser(ctx, &pb.LoginUserRequest{
		Username: param.Username,
		Password: password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErrGRPCCall(err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.ResponseSuccess(reply))
	return
}
