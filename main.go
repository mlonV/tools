package main

import (
	"time"

	"github.com/mlonV/tools/loger"
)

func main() {
	log := loger.NewLoger()
	// log.Level = loger.WARNING

	for {
		log.Debug("一个 debug")
		log.Info("一个 Info")
		log.Error("一个 Error")
		log.Fatal("一个 fatal")
		log.Warning("一个 warning")

		time.Sleep(time.Second * 2)
	}
}
