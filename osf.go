package ahelper

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Osf_an struct{}

func newOsf() Osf_an {
	return Osf_an{}
}
func (o Osf_an) Printfln(format string, a ...interface{}) {
	fmt.Printf(format+"\n", a...)

}

// 取运行目录
func (o Osf_an) Getpath_exe() string {
	path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	dir := filepath.Dir(path)
	//fmt.Println(path) // for example /home/user/main
	//fmt.Println(dir)  // for example /home/user
	return dir
}

func (o Osf_an) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 去除html标签
func (o Osf_an) Htmltrim(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile(`<[\S\s]+?>`)
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile(`<style[\S\s]+?</style>`)
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile(`<script[\S\s]+?</script>`)
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile(`<[\S\s]+?>`)
	src = re.ReplaceAllString(src, "")
	/*
	   src = re.ReplaceAllString(src, "\n")
	   //去除连续的换行符
	   re, _ = regexp.Compile("\\s{2,}")
	   src = re.ReplaceAllString(src, "\n")
	*/
	return strings.TrimSpace(src)
}


