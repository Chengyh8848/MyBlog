package datatype

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrIt(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{5, "5"},
		{interface{}(5), "5"},
		{[]int{1, 2}, "[1,2]"},
	}
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(StrIt(test.input), test.expected)
	}
}

func TestAnyBlank(t *testing.T) {
	tests := []struct {
		input    []interface{}
		expected bool
	}{
		{[]interface{}{nil, 1, 2}, true},
		{[]interface{}{"nil", 1, 2}, false},
		{[]interface{}{0, 1, 2}, true},
		{[]interface{}{"", 1, 2}, true},
		{[]interface{}{[]interface{}{}, 1, 2}, true},
	}
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(AnyBlank(test.input...), test.expected)
	}
}

func TestAllBlank(t *testing.T) {
	tests := []struct {
		input    []interface{}
		expected bool
	}{
		{[]interface{}{nil, "", 0, 0.0, []int{}}, true},
		{[]interface{}{nil, "", 0, 0.0, []int{0}}, false},
		{[]interface{}{nil, "", 0, 0.0, 1}, false},
	}
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(AllBlank(test.input...), test.expected)
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
	}{
		{" ", false},
		{"0", false},
		{"", true},
		{nil, true},
		{0, true},
		{interface{}(0), true},
		{[]int{}, true},
	}
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(IsBlank(test.input), test.expected)
	}
}

func TestStrArr2Uint64(t *testing.T) {
	assert.Equal(t, StrArr2Uint64([]string{"1", "2", "1.2", "abc"}), []uint64{1, 2, 0, 0})
}

func TestUint64Arr2Str(t *testing.T) {
	assert.Equal(t, Uint64Arr2Str([]uint64{1, 2, 3, 0}), []string{"1", "2", "3", "0"})
}

func TestUint64(t *testing.T) {
	tests := []struct {
		input    string
		expected uint64
	}{
		{"", uint64(0)},
		{"0", uint64(0)},
		{"12a", uint64(0)},
		{"12", uint64(12)},
	}
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(Uint64(test.input), test.expected)
	}
}

func TestUint32(t *testing.T) {
	tests := []struct {
		input    string
		expected uint32
	}{
		{"", uint32(0)},
		{"0", uint32(0)},
		{"12a", uint32(0)},
		{"12", uint32(12)},
	}
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(Uint32(test.input), test.expected)
	}
}

func TestInt32(t *testing.T) {
	tests := []struct {
		input    string
		expected int32
	}{
		{"", int32(0)},
		{"0", int32(0)},
		{"12a", int32(0)},
		{"12", int32(12)},
	}
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(Int32(test.input), test.expected)
	}
}

func TestInSlice(t *testing.T) {
	tests := []struct {
		input    int
		intput2  []int
		expected bool
	}{
		{1, []int{1, 2, 3}, true},
		{1, []int{}, false},
		{0, []int{}, false},
		{0, []int{0}, true},
	}
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(InSlice(test.input, test.intput2), test.expected)
	}
}

func TestInStrs(t *testing.T) {
	tests := []struct {
		input    string
		intput2  []string
		expected bool
	}{
		{"a", []string{"a", "b"}, true},
		{"ab", []string{"ab", "1"}, true},
		{"ab", []string{"abc", "1"}, false},
		{"1", []string{"a", "b"}, false},
		{"", []string{"a", "b", "0"}, false},
	}
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(InSlice(test.input, test.intput2), test.expected)
	}

}

func TestUniqueStringSlice(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"a", "b", "b", "c"}, []string{"a", "b", "c"}},
		{[]string{"a", "a", "b", "a", "b", "c", "c"}, []string{"a", "b", "c"}},
		{[]string{"c", "a", "b", "a", "b", "c", "c"}, []string{"c", "a", "b"}},
		{[]string{"a", "", "", "c", ""}, []string{"a", "", "c"}},
	}
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(UniqueSlice(test.input), test.expected)
	}
}

func TestIsNil(t *testing.T) {
	var name *string
	assert.Nil(t, name)
}

func TestIsSlice(t *testing.T) {
	assert.Equal(t, IsSlice([]string{}), true)
	assert.Equal(t, IsSlice([1]string{}), true)
	assert.Equal(t, IsSlice(1), false)

	data := []string{"a", "b", "c"}

	a(data)
}

func a(a interface{}) {
	println(reflect.ValueOf(a).Len())
}

func TestXxx(t *testing.T) {
	a := map[uint32]map[uint32]bool{}
	if _, ok := a[1][1]; !ok {
		a[1] = map[uint32]bool{}
		a[1][1] = true
	}
	println(a[1][1])
}

func TestFilteEmpty(t *testing.T) {
	assert.Equal(t, FilterEmpty([]string{"a", "", "b"}), []string{"a", "b"})
	assert.Equal(t, FilterEmpty([]int{0, 1, 2}), []int{1, 2})
}

func TestIsEmpty(t *testing.T) {
	const (
		a = 1 << 0

		b = 1 << 1
		c = 1 << 2
		d = 1 << 3
	)
	fmt.Printf("a=%d %b,b=%d %b,c=%d %b,d=%d %b\n", a, a, b, b, c, c, d, d)

	sum := a | b | c | d
	fmt.Printf("sum=%d %b\n", sum, sum)
}
