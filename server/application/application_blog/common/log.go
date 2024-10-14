package common

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"os"
	"time"
)

const ContextCenterId = "centerid"
const ContextReqUUid = "uuid"
const ContextReqSource = "source"
const ContextMethod = "method"
const ContextCcnum = "ccnum"

type ReqLogRecord struct {
	CenterId  string  `json:"centerId"`  //监控中心ID
	Uuid      string  `json:"uuid"`      //跟踪id
	Source    string  `json:"source"`    //请求来源
	Ccnum     int     `json:"ccnum"`     //调用序号
	Timestamp int64   `json:"timestamp"` //请求发起时间戳
	ReqData   []byte  `json:"reqData"`   //请求数据
	SpendTime float64 `json:"spendTime"` //该次请求所花费时间秒

	Url string

	Resp []byte `json:"resp"` //返回结果
	Err  interface{}
}

type RpcReqLogRecord struct {
	ReqLogRecord
}

func (req RpcReqLogRecord) RequestFormate() string {
	return fmt.Sprintf("uuid:[%s] ccnm:[%d] source:[%s] centerid:[%s] url:[%s] data:[%s]",
		req.Uuid, req.Ccnum, req.Source, req.CenterId, req.Url, string(req.ReqData))
}

func (req RpcReqLogRecord) ResponseFormate() string {
	respStr := "..."
	if req.Resp != nil {
		respStr = string(req.Resp)
	}

	return fmt.Sprintf("uuid:[%s] ccnum:[%d] url:[%s] spend:[%fs]  resp:[%s]  err:[%v] ", req.Uuid, req.Ccnum, req.Url, req.SpendTime, respStr, req.Err)
}

type LoggerConfig struct {
	Filename   string // 日志文件名称
	LogLevel   string // 日志级别
	MaxSize    int    // 日志大小
	MaxAge     int    // 日志保存天数
	MaxBackups int    // 日志备份个数
	LocalTime  bool   // 日志备份是否按照本地时间进行重命名 默认true
	Compress   bool   // 日志备份是否打包 默认false
	Caller     bool   // 是否开启堆栈跟踪 默认false
	Develop    bool   // 是否开启开发者模式 默认false，如果要开启需先开启堆栈跟踪
	StdOut     bool   // 是否控制台显示
}

type Logger struct {
	Log *zap.Logger
}

// 性能更好
var Log *Logger

func (l *Logger) ErrorMsg(msg string, v ...interface{}) {
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	l.Log.Error(msg)
}

func (l *Logger) InfoMsg(msg string, v ...interface{}) {
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	l.Log.Info(msg)
}

func (l *Logger) DebugMsg(msg string, v ...interface{}) {
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	l.Log.Debug(msg)
}

// 性能稍差，但是有格式宽松的结构化log，如nfow，Info，Infof；较少使用；仅使用在调用不频繁的地方
var SugarLogger *zap.SugaredLogger

func InitLogger(config LoggerConfig) {
	Log = new(Logger)
	// 获取日志文件句柄
	writer := getLogWriter(config)
	// 设置日志编码
	encoder := getEncoder(config)
	var core zapcore.Core
	switch config.LogLevel {
	case "debug":
		core = zapcore.NewCore(encoder, writer, zapcore.DebugLevel)
	case "info":
		core = zapcore.NewCore(encoder, writer, zapcore.InfoLevel)
	case "error":
		core = zapcore.NewCore(encoder, writer, zapcore.ErrorLevel)
	default:
		core = zapcore.NewCore(encoder, writer, zapcore.InfoLevel)
	}

	// 开启堆栈
	if config.Caller {
		caller := zap.AddCaller()
		// 开启开发者模式
		if config.Develop {
			development := zap.Development()
			Log.Log = zap.New(core, caller, development)
		} else {
			Log.Log = zap.New(core, caller)
		}
	} else {
		Log.Log = zap.New(core)
	}
	SugarLogger = Log.Log.Sugar()
	defer SugarLogger.Sync()
}

func getLogWriter(config LoggerConfig) zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize,
		MaxAge:     config.MaxAge,
		MaxBackups: config.MaxBackups,
		LocalTime:  config.LocalTime,
		Compress:   config.Compress,
	}
	if config.StdOut {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberjackLogger))
	} else {
		return zapcore.AddSync(lumberjackLogger)
	}
}

func getEncoder(config LoggerConfig) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	if config.Develop {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	}
	// 设置日式时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (respErr error) {
	uuid := fmt.Sprintf("UUIDNotSet_%d", time.Now().UnixNano())
	centerId := ""
	source := ""
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		if arr := md.Get(ContextReqUUid); len(arr) > 0 {
			uuid = arr[0]
		}

		if arr := md.Get(ContextReqSource); len(arr) > 0 {
			source = arr[0]
		}
	}

	ccnum := ccnumCreator.GenCcnum(uuid)
	start := time.Now()
	reqData, _ := json.Marshal(req)

	reqLogRecord := &RpcReqLogRecord{
		ReqLogRecord: ReqLogRecord{
			Uuid:      uuid,
			Ccnum:     ccnum,
			Source:    source,
			CenterId:  centerId,
			ReqData:   reqData,
			Timestamp: start.UnixNano(),
			Url:       method,
		},
	}
	Log.InfoMsg("[gRPCRequest] %s", reqLogRecord.RequestFormate())
	defer func() {
		reqLogRecord.SpendTime = time.Since(start).Seconds()
		if Log.Log.Level() == zap.DebugLevel { //输出grpc返回值
			respData, _ := json.Marshal(reply)
			reqLogRecord.Resp = respData
		}
		if respErr == nil {
			Log.InfoMsg("[gRPCRepsponse] %s", reqLogRecord.ResponseFormate())
		} else {
			Log.InfoMsg("[gRPCRepsponse] %s", reqLogRecord.ResponseFormate())
		}
	}()

	respErr = invoker(ctx, method, req, reply, cc, opts...) // invoking RPC method

	return respErr
}
