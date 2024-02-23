package boot

import (
	"log"

	"github.com/ccb1900/gocommon/config"
	"github.com/ccb1900/gocommon/env"
	"github.com/ccb1900/gocommon/logger"
)

func Boot(appname, stage string) {
	if err := env.InitEnv(stage); err != nil {
		log.Println("load env", err)
	}
	config.Init(appname)
	logger.Init()
}
