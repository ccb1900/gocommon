package boot

import (
	"log"

	"gocommon/config"
	"gocommon/env"
	"gocommon/logger"
)

func Boot(appname, stage string) {
	if err := env.InitEnv(stage); err != nil {
		log.Println("load env", err)
	}
	config.Init(appname)
	logger.Init()
}
