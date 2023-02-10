package time

import (
	"strconv"
	"time"
)

func Stamp_13() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10)
}
func Stamp_10() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
