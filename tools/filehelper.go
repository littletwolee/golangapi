package tools

import (
	"github.com/astaxie/beego"
//	"bufio"  //缓存IO
//	"fmt"
	"io/ioutil" //io 工具包
//	"io"
//	"os"
//	"log"
)

var (
	cache         string = beego.AppConfig.String("filecache")
)
type Filehelper struct{}

// func checkDirectoryIsExist(directoryname string, iscreate bool) bool {
// 	var exist = true
// 	if _ err := os.Stat(cache + "/" + directoryname); os.IsNotExist(err) {
// 		exist = false
// 	}
// 	return exist
// }

// func checkFileIsExist(filename string) (bool) {
// 	var exist = true;
// 	if _, err := os.Stat(filename); os.IsNotExist(err) {
// 		exist = false;
// 	}
// 	return exist;
// }

func (f *Filehelper)WriteFile(filename string, file []byte) (string, error) {
	path := cache + "/" + filename
	// if checkFileIsExist(path) {
	// 	return "", (&ResultHelp{}).NewErr("file is exists")
	// }
	err := ioutil.WriteFile(path, file, 0666)  //写入文件(字节数组)
	return filename, err
}

func (f *Filehelper)ReadFile(filename string) ([]byte, error){
	path := cache + "/" + filename
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
