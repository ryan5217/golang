package filetool

import (
	"os"
	"io"
	"io/ioutil"
)

// 文件接口
type Files interface {
	CheckFileExist(filename string) bool
	ReadFile(filename string) ([]byte, error)
	WriteFile(filename string, context string) (bool, error)
}

// 实现
type FileOperate struct {
	Filename string
}


// 判断文件是否存在 存在返回 true  不存在返回false
// param filename string 文件名
// return bool
func (f *FileOperate) CheckFileExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// 读文件
// param file string 文件名
// return byte, error
func (f *FileOperate) ReadFile(filename string) ([]byte, error) {
	em := []byte{}
	em, err := ioutil.ReadFile(filename)
	if err != nil {
		return em, err
	}
	return em, nil
}

// 写入文件
// param filename string 文件名
// param context  string 内容
// return bool
func (f *FileOperate) WriteFile(filename string, context string) (bool, error) {
	var file *os.File
	var err error
	if f.CheckFileExist(filename) {
		file, err = os.OpenFile(filename, os.O_APPEND, 0666) // 打开文件
		if err != nil {
			return false, err
		}
	} else {
		file, err = os.Create(filename)
		if err != nil {
			return false, err
		}
	}

	if err != nil {
		return false, err
	}
	_, err = io.WriteString(file, context)
	if err != nil {
		return false, err
	}
	return true, nil
}



