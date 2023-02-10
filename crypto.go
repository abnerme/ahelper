package ahelper

import (
	"github.com/golang-module/dongle"

)

type Crypto struct{
	
}

func newCrypto() Crypto{

	return Crypto{}
}
func (c Crypto) Md5(str string) string {
	return dongle.Encrypt.FromString(str).ByMd5().ToHexString()
}

func (c Crypto) Md5_bin(b []byte) string {
	return dongle.Encrypt.FromBytes (b).ByMd5().ToHexString()
}
