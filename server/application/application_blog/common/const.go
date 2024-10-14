package common

const (
	StatusUP      int32 = 1001 //事件类型：设备上线
	StatusDown    int32 = 1002 //事件类型：设备离线
	VideoCutOff   int32 = 1003 //事件类型：录像中断
	StorageNotice int32 = 1004 //事件类型：存储告警
)

const (
	StatusOn  int32 = 1 //设备在线
	StatusOff int32 = 2 //设备离线
)
