# tools
	log := loger.NewLoger(
		&loger.Loger{
			ToFile: true,
			FileLoger: loger.FileLoger{
				FileName:    "loger.log",
				FilePath:    "./logs/",
				FileMaxSize: 1024,
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