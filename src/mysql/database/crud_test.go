package database

import (
	"goprogramming/src/mysql/config"
	"github.com/zssky/log"
	"testing"
)

func TestInsertData(t *testing.T) {
	info := &UserInfo{
		Name: "mengdake",
		Age: 26,
	}
	// defer DB.Close()
	if err := config.InitConfig(); err != nil {
		log.Error(err)
		return
	}

	if err := InitDbConn(); err != nil {
		log.Error(err)
		return
	}
	if err := InsertData(info); err != nil{
		return
	}
}

func TestSelectData(t *testing.T) {
	if err := config.InitConfig(); err != nil {
		log.Error(err)
		return
	}

	if err := InitDbConn(); err != nil {
		log.Error(err)
		return
	}
	if err := SelectData(); err != nil{
		return
	}
}

// 测试更新数据
func TestUpdataData(t *testing.T) {
	if err := config.InitConfig(); err != nil {
		log.Error(err)
		return
	}

	if err := InitDbConn(); err != nil {
		log.Error(err)
		return
	}
	if err := UpdataData(); err != nil{
		return
	}
}