package database

import (
	"testing"

	"github.com/zssky/log"
	"goprogramming/src/mysql/config"
)

// 测试数据库长连接sql.DB
func TestInitDbConn(t *testing.T) {
	if err := config.InitConfig(); err != nil {
		log.Error(err)
		return
	}

	if err:= InitDbConn(); err != nil {
		log.Error(err)
		return
	}
	if DB == nil {
		log.Error("DB初始化失败")
		return
	}else {
		log.Info("DB初始化成功")
		DB.Close()
	}
}