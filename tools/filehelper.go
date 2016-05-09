package tools

import (
	"github.com/astaxie/beego"
//	"bufio"  //缓存IO
//	"fmt"
	"io/ioutil" //io 工具包
//	"io"
	"os"
)

var (
	cache         string = beego.AppConfig.String("filecache")
)
type Filehelper struct{}

func checkFileIsExist(filename string) (bool) {
	var exist = true;
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false;
	}
	return exist;
}

func (f *Filehelper)writefile(filename string, file []byte) error {
	path := cache + "/" + filename
	if checkFileIsExist(path) {
		return newerr("file is exists")
	}
	err := ioutil.WriteFile(path, file, 0666)  //写入文件(字节数组)
	return err
}

func (f *Filehelper)readfile(filename string) (file []byte, err error){
	file, err = ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}
