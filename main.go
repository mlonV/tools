package main

import (
	"time"

	"github.com/mlonV/tools/loger"
)

func main() {
	// log := loger.NewLoger()
	// log.Level = loger.WARNING
	// log.ToFile = true

	// 分割线
	log := loger.NewLoger()
	fileLogObj := loger.FileLoger{
		FileName:    "loger.log",
		FilePath:    "/Users/apple/Desktop/go_dev/tools/test",
		FileMaxSize: 10 * 1024,
	}
	log.FileLoger = fileLogObj
	log.ToFile = true

	for {
		id := 100
		name := "test"
		log.Debug("一个 debug id = %d , name = %s", id, name)
		log.Info("一个 Info")
		log.Error("一个 Error id = %d , name = %s", id, name)
		log.Fatal("一个 fatal")
		log.Warning("一个 warning")

		time.Sleep(time.Second * 2)
	}
}
