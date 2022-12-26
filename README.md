# tools
	// 分割线
	log := loger.NewLoger(
		&loger.Loger{
			ToFile:          true,
			WithFuncAndFile: false,
			Level:           2,
			FileLoger: loger.FileLoger{
				FileName:    "loger.log",
				FilePath:    "./logs/",
				FileMaxSize: 2048,
				FileSaveNum: 5,
			},
		},
	)

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