package errex

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrexI interface {
	Error() string

	GetCode() uint32
	GetDesc() string
	GrpcError() error
}

type Errex struct {
	Code    uint32
	Message string
	Desc    string
}

func (e Errex) Error() string {
	return fmt.Sprintf("code:%d desc:%s", e.Code, e.Desc)
}
func (e Errex) GetCode() uint32 {
	return e.Code
}
func (e Errex) GetDesc() string {
	return e.Desc
}

func (e Errex) GrpcError() error {
	return status.Errorf(codes.Code(e.Code), e.Message+" "+e.Desc)
}

const SERVER_COMMON_ERROR uint32 = 100001
const REUQEST_PARAM_ERROR uint32 = 100002
const TOKEN_EXPIRE_ERROR uint32 = 100003
const TOKEN_GENERATE_ERROR uint32 = 100004

const DB_COMMON_ERROR uint32 = 100005
const DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100006
const DB_CONN_NOT_EXIST = 100007
const DB_RECORD_NOT_FOUND = 100008

const DA_HTTPSERVER_ERROR = 100009
const LICENCE_EXCEED_LIMIT = 100010 //licence key认证数量超出限制

func NewCommonError(format string, a ...interface{}) *Errex {
	return &Errex{SERVER_COMMON_ERROR, "内部错误", fmt.Sprintf(format, a...)}
}

func NewParamError(format string, a ...interface{}) *Errex {
	return &Errex{REUQEST_PARAM_ERROR, "参数错误", fmt.Sprintf(format, a...)}
}

func NewDbCommonError(format string, a ...interface{}) *Errex {
	return &Errex{DB_COMMON_ERROR, "数据库错误", fmt.Sprintf(format, a...)}
}

func NewDbConnNotExistError(format string, a ...interface{}) *Errex {
	return &Errex{DB_CONN_NOT_EXIST, "数据库连接不存在", fmt.Sprintf(format, a...)}
}

func NewRecordNotFoundError(format string, a ...interface{}) *Errex {
	return &Errex{DB_RECORD_NOT_FOUND, "记录不存在", fmt.Sprintf(format, a...)}
}

func NewPlatformError(format string, a ...interface{}) *Errex {
	return &Errex{DA_HTTPSERVER_ERROR, "平台错误", fmt.Sprintf(format, a...)}
}

func NewLicenceExceedLimitError() *Errex {
	return &Errex{LICENCE_EXCEED_LIMIT, "超出授权路数限制", ""}
}
