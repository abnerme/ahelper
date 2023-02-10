package ahelper
import (
	"fmt"
	"os"
	"path/filepath"
)

type Osf_an struct{}

func newOsf() Osf_an{
	return Osf_an{}
}
func (o Osf_an) Printfln(format string, a ...interface{}) {
	fmt.Printf(format+"\n", a...)

}

// 取运行目录
func (o Osf_an) Getpath_exe() string{
	path, err := os.Executable()
	if err != nil {
    fmt.Println(err)
	}
	dir := filepath.Dir(path)
	//fmt.Println(path) // for example /home/user/main
	//fmt.Println(dir)  // for example /home/user
	return dir
}
