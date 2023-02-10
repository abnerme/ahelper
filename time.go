package ahelper

import (
	"strconv"
	"time"

	
)

type Time struct{
	
}

func newTime() Time{

	return Time{}
}

func (t Time) Stamp_13() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10)
}
func (t Time) Stamp_10() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

