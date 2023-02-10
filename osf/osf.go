package osf

import (
	"fmt"
	"os"
	"path/filepath"
)

func Printfln(format string, a ...interface{}) {
	fmt.Printf(format+"\n", a...)

}

// 取运行目录
func Getpath_exe() string{
	path, err := os.Executable()
	if err != nil {
    fmt.Println(err)
	}
	dir := filepath.Dir(path)
	//fmt.Println(path) // for example /home/user/main
	//fmt.Println(dir)  // for example /home/user
	return dir
}
