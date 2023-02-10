package ahelper

import (
	"strconv"
	"time"

	
)

type Time_an struct{
	
}

func newTime() Time_an{

	return Time_an{}
}

func (t Time_an) Stamp_13() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10)
}
func (t Time_an) Stamp_10() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

