package loger

import (
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

type FileLoger struct {
	FileName    string
	FilePath    string
	fileObj     *os.File
	FileMaxSize int64  // 每个文件的大小，轮训
	FileSaveNum uint64 // 文件保存数量，超过数量则删除最旧的，保留新的

	// 里面保存每一个创建出来的文件名，若重启则需要读取出来所有文件重新录入channel
	FC chan string
}

var (
	withonce sync.Once
)

func (f *FileLoger) WriteFileLoger() *os.File {
	withonce.Do(f.AddOldFileToFC)
	fullFilePath := path.Join(f.FilePath, f.FileName)
	//
	if f.fileObj == nil {
		f.fileObj = f.GetFileObj(fullFilePath)
	}

	// 获取当前文件大小
	fi, err := f.fileObj.Stat()
	if err != nil || fi == nil {
		// 获取不到状态直接panic
		panic(fmt.Sprintf("get file stat failed err : %v ,obj : %v", err, fi))
	}

	// 若超过指定大小，则文件重命名。
	if fi.Size() > f.FileMaxSize {
		f.RenameFile(fullFilePath)
		f.fileObj = f.GetFileObj(fullFilePath)
		return f.fileObj
	}
	return f.fileObj

}

func (f *FileLoger) GetFileObj(fullFilePath string) *os.File {
	file, err := os.OpenFile(fullFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	return file
}

// 判断是否超过指定大小，若超过则切换文件，打开新文件输入
func (f *FileLoger) RenameFile(fullFilePath string) {
	newFileName := fullFilePath + "-" + time.Now().Format("20060102150405")
	os.Rename(fullFilePath, newFileName)
	f.DeleteFile()
	f.FC <- newFileName
	f.fileObj.Close()
}

// 把老的文件加入到FC
func (f *FileLoger) AddOldFileToFC() {
	de, err := os.ReadDir(f.FilePath)
	if err != nil {
		panic(err)
	}
	for _, v := range de {
		if v.Name() == f.FileName {
			continue
		}
		oldFile := strings.Split(v.Name(), "-")[0]
		if oldFile == f.FileName {
			if len(f.FC) >= int(f.FileSaveNum) {
				f.DeleteFile()
			}
			f.FC <- path.Join(f.FilePath, v.Name())

		}
	}
}

// 获取文件的绝对路径
func (f *FileLoger) GetFullFilePath() string {
	return path.Join(f.FilePath, f.FileName)
}

func (f *FileLoger) DeleteFile() {
	if len(f.FC) >= int(f.FileSaveNum) {
		file := <-f.FC
		os.Remove(file)
	}

}
