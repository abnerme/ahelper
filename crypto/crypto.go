package crypto

import (
	"github.com/golang-module/dongle"
)

func Md5(str string) string {
	return dongle.Encrypt.FromString(str).ByMd5().ToHexString()
}

func Md5_bin(b []byte) string {
	return dongle.Encrypt.FromBytes (b).ByMd5().ToHexString()
}
