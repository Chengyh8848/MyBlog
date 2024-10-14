package interceptor

import (
	"context"

	"domain_blog/conf"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TranslateCtxInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("context 为空")
	}

	if len(md.Get(conf.ContextReqUUid)) > 0 {
		ctx = context.WithValue(ctx, conf.ContextReqUUid, md.Get(conf.ContextReqUUid)[0])
	}

	return handler(ctx, req)
}
