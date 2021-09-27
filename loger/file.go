package loger

import (
	"os"
	"path"
)

type FileLoger struct {
	FileName    string
	FilePath    string
	fileObj     *os.File
	FileMaxSize int
}

func (f *FileLoger) NewFileLoger() {
	f.initFile()
}

func (f *FileLoger) initFile() {
	fullFilePath := path.Join(f.FilePath, f.FileName)
	file, err := os.OpenFile(fullFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	f.fileObj = file
}
