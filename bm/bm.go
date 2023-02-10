package bm

import (
	"bytes"
	"io"
	"net/url"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"github.com/axgle/mahonia"
)

func Urlencode(u string) string{
	return url.QueryEscape(u)
}

func Urldecode(u string) string{
	str,_:= url.QueryUnescape(u)
	return str
}

//gbk转utf8

func GbkToUtf8(s []byte) ([]byte, error) {

	reader :=transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())	
	d, e :=io.ReadAll(reader)
	if e !=nil {	
	return nil, e	
	}	
	return d, nil	
}
	
	//utf8转gbk
	
func Utf8ToGbk(s []byte) ([]byte, error) {
	
	reader :=transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	
	d, e :=io.ReadAll(reader)
	
	if e !=nil {
	
	return nil, e
	
	}
	
	return d, nil
	
}

func ConvertToString(src string, srcCode string, tagCode string) string {
//ConvertToString("襄阳","gbk","ascii")
	srcCoder := mahonia.NewDecoder(srcCode)
   
	srcResult := srcCoder.ConvertString(src)
   
	tagCoder := mahonia.NewDecoder(tagCode)
   
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
   
	result := string(cdata)
   
	return result
   
}
	
