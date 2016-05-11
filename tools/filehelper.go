package tools

import (
	"github.com/astaxie/beego"
//	"bufio" 
//	"fmt"
	"io/ioutil"
//	"io"
	"os"
//	"log"
)

var (
	filecache         string = beego.AppConfig.String("filecache")
)
type Filehelper struct{}

// func checkDirectoryIsExist(directoryname string, iscreate bool) bool {
// 	var exist = true
// 	if _ err := os.Stat(cache + "/" + directoryname); os.IsNotExist(err) {
// 		exist = false
// 	}
// 	return exist
// }

func checkFileIsExist(dirpath string, filename string) (bool) {
	var exist = true;
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		_ = os.MkdirAll(dirpath, 0777)
	}
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false;
	}
	return exist;
}

func (f *Filehelper)WriteFile(filename string, file []byte) (string, error) {
	if checkFileIsExist(filecache, filename) {
		return "", (&ResultHelp{}).NewErr("file is exists")
	}
	path := filecache + "/" + filename
	err := ioutil.WriteFile(path, file, 0777)
	return filename, err
}

func (f *Filehelper)ReadFile(filename string) ([]byte, error){
	if checkFileIsExist(filecache, filename) {
		return nil, (&ResultHelp{}).NewErr("file is exists")
	}
	path := filecache + "/" + filename
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
