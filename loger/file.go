package loger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLoger struct {
	FileName    string
	FilePath    string
	fileObj     *os.File
	FileMaxSize int64
}

func (f *FileLoger) WriteFileLoger() *os.File {
	fullFilePath := path.Join(f.FilePath, f.FileName)
	//
	var file *os.File = f.GetFileObj(fullFilePath)

	// 获取当前文件大小
	fi, err := file.Stat()
	if err != nil || fi == nil {
		// 获取不到状态直接panic
		panic(fmt.Sprintf("get file stat failed err : %v ,obj : %v", err, fi))
	}

	// 若超过指定大小，则文件重命名。
	if fi.Size() > f.FileMaxSize {
		f.RenameFile(fullFilePath)
		return f.GetFileObj(fullFilePath)
	}
	return file

}

func (f *FileLoger) GetFileObj(fullFilePath string) *os.File {
	file, err := os.OpenFile(fullFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	return file
}

// 判断是否超过指定大小，若超过则切换文件，打开新文件输入
func (f *FileLoger) RenameFile(fullFilePath string) {
	// fmt.Println(time.Now().Format("200601021504"), fullFilePath)
	os.Rename(fullFilePath, fullFilePath+"-"+time.Now().Format("20060102150405"))
}
