package archive

import (
	"application_blog/common"
	"application_blog/conf"
	client "application_blog/internal/pkg/grpc"
	pb "application_blog/internal/protocal/pb/blog"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetArchives(c *gin.Context) {
	redisKey := common.ARCHIVE_BLOG_MAP
	res, err := common.Client.GetSet(redisKey)
	if err != nil || res == "" {
		c.JSON(http.StatusOK, common.ResponseSuccess(nil))
		return
	}
	var tempMap map[string]ArchiveBlog
	err = json.Unmarshal([]byte(res), &tempMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErr("获取缓存信息报错"))
		return
	}
	if tempMap == nil || len(tempMap) == 0 {
		c.JSON(http.StatusOK, common.ResponseSuccess(nil))
		return
	}
	ctx, cancel := common.SetContextWithTimeout(c, conf.GrpcTimeOut)
	defer cancel()
	reply, err := client.GrpcClients.BlogHostClient.GetGroupYearMonthByIsPublished(ctx, &pb.GetGroupYearMonthByIsPublishedRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErr("获取文章归档报错"))
		return
	}

}
