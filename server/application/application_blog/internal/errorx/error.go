package errorx

type ErrCode int

const (
	ErrInternal ErrCode = iota + 1000
	ErrParams
	ErrParse
	ErrGrpc
	ErrNotExitUser
	ErrPassword
	ErrFileExt
	ErrUserLocked
	ErrRedis
	ErrTokeExpired
	ErrCustom
	ErrNotAllowed
	ErrNotToken
	ErrAuth2NotFound
	ErrClientIdNotFound
	ErrNotFoundLicense
	ErrParseLicense
	ErrLicense
	ErrGetMachine
	ErrUploadFile
	ErrLicenseInvalid
)

var TextErr = map[ErrCode]string{
	ErrInternal:         "服务错误",
	ErrParams:           "参数错误",
	ErrParse:            "解析错误",
	ErrGrpc:             "RPC失败",
	ErrNotExitUser:      "用户不存在",
	ErrPassword:         "用户密码错误",
	ErrFileExt:          "文件扩展类型错误",
	ErrUserLocked:       "用户已冻结",
	ErrRedis:            "组件获取失败",
	ErrTokeExpired:      "用户令牌失效",
	ErrNotToken:         "用户token不存在",
	ErrNotAllowed:       "用户不允许在该电脑登录",
	ErrAuth2NotFound:    "Oauth2授权类型不存在",
	ErrClientIdNotFound: "客户端标识不存在",
	ErrNotFoundLicense:  "系统未发现License授权文件",
	ErrParseLicense:     "系统License授权不合法",
	ErrLicense:          "系统License错误",
	ErrGetMachine:       "系统设备错误",
	ErrUploadFile:       "系统License文件上传失败",
	ErrLicenseInvalid:   "系统License文件不合法",
}

func (e ErrCode) Error() string {
	if v, ok := TextErr[e]; ok {
		return v
	}
	return "不存在错误"
}

func (e ErrCode) String() string {
	if v, ok := TextErr[e]; ok {
		return v
	}
	return "不存在错误"
}
