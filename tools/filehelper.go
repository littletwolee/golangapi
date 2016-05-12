package tools

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"os"
	"golangapi/models"
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

func (f *Filehelper)UploadFileToMongo(filemode models.Filemodel) (string, error) {
	var filename string
	mongogridfshelper := &MongoGridFSHelper{}
	redishelper := &RedisHelper{}
	if filemode.Filename == "" {
		filename = GetGuid() 
	} else {
		filename = filemode.Filename
	}
	switch {
	case filemode.Currentchunk == 0 && filemode.Currentchunk == filemode.Maxchunks - 1 :
		return mongogridfshelper.UploadFile(filemode)
	case filemode.Currentchunk < filemode.Maxchunks:
		if filemode.Currentchunk != 0 {
			if filechunkdata ,err := redishelper.GetVByK(filename, "bytes"); err == nil {
				filemode.Filedata = append(filechunkdata.([]byte), filemode.Filedata...)
			}
		}
		if filemode.Currentchunk == filemode.Maxchunks - 1 {
			return mongogridfshelper.UploadFile(filemode)
		} else {
			if err := redishelper.SetKVBySETEX(filename, filemode.Filedata, 60); err == nil {
				return filename, nil
			}
		}
	}
	return "", (&ResultHelp{}).NewErr("server err")
}
