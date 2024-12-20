package about

import (
	"application_blog/common"
	"application_blog/conf"
	client "application_blog/internal/pkg/grpc"
	pb "application_blog/internal/protocal/pb/blog"
	"github.com/gin-gonic/gin"
	"net/http"
)

func About(c *gin.Context) {
	ctx, cancel := common.SetContextWithTimeout(c, conf.GrpcTimeOut)
	defer cancel()
	reply, err := client.GrpcClients.BlogHostClient.GetAboutInfo(ctx, &pb.GetAboutInfoRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErrGRPCCall(err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.ResponseSuccess(reply))
	return
}
