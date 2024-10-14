package interceptor

import (
	"context"
	"domain_blog/common"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"time"

	"domain_blog/conf"

	"google.golang.org/grpc"
)

func LoggerInterceptor(l *common.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func(begin time.Time) {
			var panicStack = []byte{}
			if e := recover(); e != nil {
				panicStack = debug.Stack()
				err = fmt.Errorf("异常:%v", e)

				l.ErrorMsg("异常:%v %s", e, string(panicStack))
			}
			if err == nil {
				l.InfoMsg("gRPC 调用响应:[%s] [%s] 耗时:%s Err:%v %s", ctx.Value(conf.ContextReqUUid), info.FullMethod, time.Since(begin), err, string(panicStack))
			} else {
				l.ErrorMsg("gRPC 调用响应:[%s] [%s] 耗时:%s Err:%v %s", ctx.Value(conf.ContextReqUUid), info.FullMethod, time.Since(begin), err, string(panicStack))
			}
		}(time.Now())
		//当数据库连接数大于等于设置连接数时记录日志 影响并发
		// if conn, err := pool.Pools.GetDBConn(ctx); err == nil {
		// 	if db, err := conn.DB(); err == nil {
		// 		s := db.Stats()
		// 		if s.MaxOpenConnections != 0 && s.OpenConnections >= s.MaxOpenConnections && s.InUse >= s.MaxOpenConnections {
		// 			l.Info("数据连接状态:[%s] [%s] [%s]", ctx.Value(conf.ContextReqUUid), ctx.Value(conf.ContextCenterId), datatype.StrIt(s))
		// 		}
		// 	}
		// }

		reqData, _ := json.Marshal(req)
		l.InfoMsg("gRPC 调用请求: [%s] [%s] with body [%s] ", ctx.Value(conf.ContextReqUUid), info.FullMethod, string(reqData))
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("客户端主动断开连接")
		default:
			return handler(ctx, req)
		}
	}
}
