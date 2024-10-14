package middleware

import (
	"application_blog/common"
	"application_blog/conf"
	client "application_blog/internal/pkg/grpc"
	pb "application_blog/internal/protocal/pb/blog"
	"github.com/gin-gonic/gin"
)

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		ctx, cancel := common.SetContextWithTimeout(c, conf.GrpcTimeOut)
		defer cancel()
		_, err := client.GrpcClients.BlogHostClient.AuthToken(ctx, &pb.AuthTokenRequest{
			Token: token,
		})
		if err != nil {
			c.Abort()
			c.JSON(401, common.NewErrInvalidToken(err.Error()))
			return
		} else {
			c.Next()
		}
	}
}
