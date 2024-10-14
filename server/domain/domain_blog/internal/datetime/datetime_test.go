package datetime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeFormat(t *testing.T) {
	tt, _ := time.Parse(TIME_TEMPLATE, "2011-02-03 04:05:06")
	assert.Equal(t, TimeFormat(tt), "2011-02-03 04:05:06")
}

func TestParseTime(t *testing.T) {
	tt := time.Date(2022, time.Month(12), 24, 15, 36, 28, 0, time.Local)
	assert.Equal(t, tt, *ParseTime("2022-12-24 15:36:28"))
}
