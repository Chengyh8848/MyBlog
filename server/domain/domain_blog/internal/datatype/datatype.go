package datatype

import (
	"encoding/json"
	"reflect"
	"strconv"
)

// interface 转 string
func StrIt(value interface{}) string {
	var key string
	if value == nil {
		return ""
	}

	switch v := value.(type) {
	case string:
		return v
	case float64:
		key = strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		key = strconv.FormatFloat(float64(v), 'f', -1, 64)
	case int:
		key = strconv.Itoa(v)
	case uint:
		key = strconv.Itoa(int(v))
	case int8:
		key = strconv.Itoa(int(v))
	case uint8:

		key = strconv.Itoa(int(v))
	case int16:

		key = strconv.Itoa(int(v))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		key = strconv.Itoa(int(v))
	case uint32:
		key = strconv.Itoa(int(v))
	case int64:
		key = strconv.FormatInt(v, 10)
	case uint64:
		key = strconv.FormatUint(v, 10)
	case []byte:
		key = string(v)
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

func AnyBlank(values ...interface{}) bool {
	for _, v := range values {
		if IsBlank(v) {
			return true
		}
	}
	return false
}

func AllBlank(values ...interface{}) bool {
	for _, v := range values {
		if !IsBlank(v) {
			return false
		}
	}
	return true
}

// 判断一个值是否为空值
func IsBlank(data interface{}) bool {
	if data == nil {
		return true
	}
	value := reflect.ValueOf(data)
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Slice:
		return value.Len() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func Marshal(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func MarshalToByte(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}

func StrArr2Uint64(ids []string) (uint64Arr []uint64) {
	for _, id := range ids {
		uint64Id, _ := strconv.ParseUint(id, 10, 64)
		uint64Arr = append(uint64Arr, uint64Id)
	}
	return uint64Arr
}

func Uint64Arr2Str(ids []uint64) (strArr []string) {
	for _, id := range ids {
		strArr = append(strArr, strconv.FormatUint(id, 10))
	}
	return strArr
}

func Uint64(strId string) uint64 {
	id, _ := strconv.ParseUint(strId, 10, 64)
	return id
}

func Uint32(num string) uint32 {
	n, _ := strconv.ParseUint(num, 10, 32)
	return uint32(n)
}

func Int32(num string) int32 {
	n, _ := strconv.ParseInt(num, 10, 32)
	return int32(n)
}

func FilterEmpty[T Alphanumeric](slice []T) []T {
	var zero T
	var res []T
	for _, num := range slice {
		if num != zero {
			res = append(res, num)
		}
	}
	return res
}

func InSlice[T Alphanumeric](id T, slice []T) bool {
	for _, v := range slice {
		if v == id {
			return true
		}
	}
	return false
}

func UniqueSlice[T Alphanumeric](slice []T) []T {
	m := map[T]int{}
	uniqSlice := []T{}
	for _, s := range slice {
		m[s] += 1
		if m[s] == 1 {
			uniqSlice = append(uniqSlice, s)
		}
	}
	return uniqSlice
}

// 判断一个接口是否是列表类型
func IsSlice(s interface{}) bool {
	t := reflect.TypeOf(s)
	return t.Kind() == reflect.Array || t.Kind() == reflect.Slice
}

type Int interface {
	int | int8 | int16 | int32 | int64
}

type Uint interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float32 | float64
}

type Number interface {
	Int | Uint | Float
}

type Alphanumeric interface {
	Number | string
}
