package main

import (
	"log"

	"shortdesc-api/internal/config"
	"shortdesc-api/internal/handler"
	"shortdesc-api/internal/mediawiki"
)

func main() {
	err := config.Load("env.yaml")
	if err != nil {
		log.Panic(err)
	}

	var mwInstanceConfig config.MWInstanceConfig

	for _, cfg := range config.Env.MWInstanceConfigs {
		if cfg.ID == config.Env.MainMWInstance {
			mwInstanceConfig = cfg
		}
	}

	err = mediawiki.Init(mwInstanceConfig)
	if err != nil {
		log.Panic(err)
	}

	handler.HandleRequests()
}
