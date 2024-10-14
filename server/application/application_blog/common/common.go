package common

import (
	pb "application_blog/internal/protocal/pb/blog"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"math"
	"reflect"
	"time"
)

var (
	ErrInvalidParam  = &Error{Code: 1001, Message: "当前参数错误"}
	ErrInvalidToken  = &Error{Code: 1002, Message: "token错误或过期"}
	ErrGetDateFailed = &Error{Code: 1005, Message: "数据库操作失败"}
	ErrFuncHandle    = &Error{Code: 1006, Message: "功能处理错误"}
	ErrGRPCCall      = &Error{Code: 1007, Message: "GRPC请求失败"}
	TimeLayout       = "2006-01-02 15:04:05"
)

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (e *Error) SetError(err error) *Error {
	e.Data = err.Error()
	return e
}

func (e *Error) SetInfo(s string) *Error {
	e.Data = s
	return e
}
func (e *Error) AddMessage(msg string) *Error {
	e.Message += ":" + msg
	return e
}

func NewErrInvalidParam(msg string, a ...interface{}) *Error {
	return &Error{Code: 1001, Message: fmt.Sprintf("当前参数错误:"+msg, a...)}
}

func NewErrInvalidToken(msg string, a ...interface{}) *Error {
	return &Error{Code: 1002, Message: fmt.Sprintf("token错误或过期:"+msg, a...)}
}

func NewErrWeekPassword() *Error {
	return &Error{Code: 1004, Message: "密码不符合规则请重试"}
}
func NewErrGetDateFailed(msg string, a ...interface{}) *Error {
	return &Error{Code: 1005, Message: fmt.Sprintf("数据库操作失败:"+msg, a...)}
}

func NewErrFuncHandle(msg string, a ...interface{}) *Error {
	return &Error{Code: 1006, Message: fmt.Sprintf("功能处理错误:"+msg, a...)}
}

func NewErrGRPCCall(msg string, a ...interface{}) *Error {
	return &Error{Code: 1007, Message: fmt.Sprintf("GRPC请求失败:"+msg, a...)}
}

func NewPlatformErr(msg string, a ...interface{}) *Error {
	return &Error{Code: 1008, Message: fmt.Sprintf("平台请求失败:"+msg, a...)}
}

func NewDataErr(msg string, a ...interface{}) *Error {
	return &Error{Code: 1009, Message: fmt.Sprintf("数据错误:"+msg, a...)}
}

func NewErrLicense(msg string, a ...interface{}) *Error {
	return &Error{Code: 1010, Message: fmt.Sprintf("授权码过期或者错误:"+msg, a...)}
}

func NewErr(msg string) *Error {
	return &Error{Code: 1011, Message: msg}
}

type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseSuccess(data interface{}) HttpResponse {
	return HttpResponse{
		Code:    0,
		Message: "success",
		Data:    data,
	}
}

func ResponseError(code int, message string, a ...interface{}) HttpResponse {
	return HttpResponse{
		Code:    code,
		Message: fmt.Sprintf(message, a...),
	}
}

func SetContextWithTimeout(c *gin.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	md := metadata.MD{}

	//sourceName := conf.Cfg.Server.Name

	//md.Set(conf.ContextSourceName, sourceName)
	//md.Set(conf.ContextReqUUid, logger.GetUuidFromGinCtx(c))

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	ctx = metadata.NewOutgoingContext(ctx, md)

	return ctx, cancel
}

// BuildCondition ...
func BuildCondition(query *[]*pb.Query, key string, value interface{}, opt pb.Operator, rel pb.Relation, options ...Option) {
	if reflect.ValueOf(value).IsZero() {
		return
	}

	q := &pb.Query{
		Key:      key,
		Value:    fmt.Sprintf("%v", value),
		Operator: opt,
		Relation: rel,
	}
	// 添加可选参数
	for _, option := range options {
		option(q)
	}

	*query = append(*query, q)
}

// Option 操作Query
type Option func(*pb.Query)

// LeftParenthesis 添加左括号
func LeftParenthesis(left uint32) Option {
	return func(q *pb.Query) {
		q.LeftParenthesis = left
	}
}

// RightParenthesis ...
func RightParenthesis(right uint32) Option {
	return func(q *pb.Query) {
		q.RightParenthesis = right
	}
}

func ParseLongTimeString(tStr string) int64 {
	timeStamp, _ := time.ParseInLocation(TimeLayout, tStr, time.Local)
	return timeStamp.Unix()
}

func ParseStringTimeLong(tLong int64) string {
	return time.Unix(tLong, 0).Format(TimeLayout)
}

func CompareDate(date1, date2 string) bool {
	timeStamp1, _ := time.ParseInLocation(TimeLayout, date1, time.Local)
	tLong1 := timeStamp1.Unix()
	timeStamp2, _ := time.ParseInLocation(TimeLayout, date2, time.Local)
	tLong2 := timeStamp2.Unix()
	if tLong1 <= tLong2 {
		return false
	} else {
		return true
	}
}

func ConvertToInt32(value interface{}) (int32, error) {
	switch v := value.(type) {
	case int:
		return int32(v), nil
	case int8:
		return int32(v), nil
	case int16:
		return int32(v), nil
	case int32:
		return v, nil
	case int64:
		if v > math.MaxInt32 || v < math.MinInt32 {
			return 0, fmt.Errorf("int64 value out of range for int32: %d", v)
		}
		return int32(v), nil
	case uint:
		if uint(math.MaxInt32) < v {
			return 0, fmt.Errorf("uint value out of range for int32: %d", v)
		}
		return int32(v), nil
	case uint8:
		return int32(v), nil
	case uint16:
		return int32(v), nil
	case uint32:
		if v > math.MaxInt32 {
			return 0, fmt.Errorf("uint32 value out of range for int32: %d", v)
		}
		return int32(v), nil
	case uint64:
		if uint64(math.MaxInt32) < v {
			return 0, fmt.Errorf("uint64 value out of range for int32: %d", v)
		}
		return int32(v), nil
	case float32:
		if v > float32(math.MaxInt32) || v < float32(math.MinInt32) {
			return 0, fmt.Errorf("float32 value out of range for int32: %f", v)
		}
		return int32(v), nil
	case float64:
		if v > math.MaxInt32 || v < math.MinInt32 {
			return 0, fmt.Errorf("float64 value out of range for int32: %f", v)
		}
		return int32(v), nil
	default:
		return 0, fmt.Errorf("unsupported type %T for conversion to int32", v)
	}
}
