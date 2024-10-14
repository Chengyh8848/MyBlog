package helper

import (
	"bytes"
	pb "domain_blog/infrastructure/grpc/pb"
	"fmt"
	"log"
	"strconv"
	"time"
	"unicode"
)

// 时间格式化标准字符串
const (
	TimeLayout  = "2006-01-02 15:04:05"
	MaxPageSize = 10000
)

// 数据库SQL关系条件
const (
	SqlGT    string = " > "    // 大于
	SqlEQUAL string = " = "    // 等于
	SqlLT    string = " < "    // 小于
	SqlNEQ   string = " != "   // 不等于
	SqlLIKE  string = " like " // 包含
	SqlGE    string = " >= "   // 大于等于
	SqlLE    string = " <= "   // 小于等于
	SqlIN    string = " in "   // in
)

// 数据库SQL逻辑条件
const (
	SqlAND string = " and " // 逻辑与
	SqlOR  string = " or "  // 逻辑或
)

var OperatorMap = map[pb.Operator]string{
	pb.Operator_GT:    SqlGT,
	pb.Operator_EQUAL: SqlEQUAL,
	pb.Operator_LT:    SqlLT,
	pb.Operator_NEQ:   SqlNEQ,
	pb.Operator_LIKE:  SqlLIKE,
	pb.Operator_GE:    SqlGE,
	pb.Operator_LE:    SqlLE,
	pb.Operator_IN:    SqlIN,
}

var RelationMap = map[pb.Relation]string{
	pb.Relation_AND: SqlAND,
	pb.Relation_OR:  SqlOR,
}

// condition 查询条件结构
type condition struct {
	Key              string // 表字段名称
	Value            string // 表字段值
	Operator         string // 关系条件
	Relation         string // 逻辑条件
	LeftParenthesis  uint32 // 左括号的个数
	RightParenthesis uint32 // 右括号的个数
}

func CheckPageLimit(limit uint32) uint32 {
	if limit > MaxPageSize {
		limit = MaxPageSize
	}
	return limit
}

func setConditions(conditions []*condition) string {
	var buffer bytes.Buffer
	var str string
	var key string
	var value interface{}

	for k, v := range conditions {
		if k > 0 {
			buffer.WriteString(v.Relation)
		}

		for i := v.LeftParenthesis; i > 0; i-- {
			buffer.WriteString(" (")
		}
		v.Key = conversion(v.Key)
		switch v.Key {
		case "created_at":
			key = "created_at"
			value = ParseTimeString(v.Value)
		case "updated_at":
			key = "updated_at"
			value = ParseTimeString(v.Value)
		default:
			key = v.Key
			value = v.Value
		}

		if v.Operator == SqlIN {
			str = fmt.Sprintf("%s%s(%v)", key, v.Operator, value)
		} else if v.Operator == SqlLIKE {
			str = fmt.Sprintf("%s%s'%%%v%%'", key, v.Operator, value)
		} else {
			str = fmt.Sprintf("%s%s'%v'", key, v.Operator, value)
		}
		buffer.WriteString(str)

		for i := v.RightParenthesis; i > 0; i-- {
			buffer.WriteString(") ")
		}
	}
	condition := buffer.String()
	if condition == "" {
		condition = "1 = 1"
	}
	return condition
}

// ToSQL 构建查询SQL语句
func ToSQL(queries []*pb.Query) (sql string) {
	var conditions []*condition
	for _, v := range queries {
		conditions = append(conditions, &condition{
			Key:              v.Key,
			Value:            v.Value,
			Operator:         OperatorMap[v.Operator],
			Relation:         RelationMap[v.Relation],
			LeftParenthesis:  v.LeftParenthesis,
			RightParenthesis: v.RightParenthesis,
		})
	}
	sql = setConditions(conditions)
	return
}

type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		_, _ = b.Write(val)
	case rune:
		_, _ = b.WriteRune(val)
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()
	_, _ = b.WriteString(s)
	return b
}

func conversion(key string) string {
	buffer := NewBuffer()
	for i, r := range key {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

func ParseTimeString(tStr string) time.Time {
	if len(tStr) == 0 {
		return time.Time{}
	}
	t, err := time.ParseInLocation(TimeLayout, tStr, time.Local)
	if err != nil {
		return time.Time{}
	}
	return t
}

func ParseLongTimeString(tStr string) int64 {
	timeStamp, _ := time.ParseInLocation(TimeLayout, tStr, time.Local)
	return timeStamp.Unix()
}

func ParseStringTimeLong(tLong int64) string {
	return time.Unix(tLong, 0).Format(TimeLayout)
}
