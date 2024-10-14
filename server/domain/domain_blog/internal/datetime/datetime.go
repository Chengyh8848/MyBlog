package datetime

import "time"

const TIME_TEMPLATE = "2006-01-02 15:04:05"

func NowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func TimeFormat(t time.Time) string {
	return t.Format(TIME_TEMPLATE)
}

func ParseTime(t string) *time.Time {
	tt, err := time.ParseInLocation(TIME_TEMPLATE, t, time.Local)
	if err != nil {
		return nil
	}
	return &tt
}

//将此种类型的时间2001-02-03T00:00:00+08:00转换为2001-02-03
func FormatDate(date *string) string {
	if date == nil {
		return ""
	}
	if len(*date) >= len("2001-02-03") {
		return (*date)[:len("2001-02-03")]
	}
	return *date
}

func FormateDateForDB(data *string) *string {
	if data == nil || *data == "" {
		return nil
	}
	return data
}
