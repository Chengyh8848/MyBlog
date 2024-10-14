package user

import (
	"application_blog/common"
	"application_blog/conf"
	client "application_blog/internal/pkg/grpc"
	pb "application_blog/internal/protocal/pb/blog"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Logout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	ctx, cancel := common.SetContextWithTimeout(c, conf.GrpcTimeOut)
	defer cancel()
	_, err := client.GrpcClients.BlogHostClient.LogoutUser(ctx, &pb.LogoutUserRequest{
		Token: token,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErrGRPCCall(err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.ResponseSuccess(nil))
	return
}
