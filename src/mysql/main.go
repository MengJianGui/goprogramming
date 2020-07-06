package main

import (
	"github.com/zssky/log"

	"goprogramming/src/mysql/config"
	"goprogramming/src/mysql/database"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Error(err)
		return
	}

	if err := database.InitDbConn(); err != nil {
		log.Error(err)
		return
	}
}
