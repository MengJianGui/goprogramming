package config

import (
	"fmt"
	"github.com/zssky/log"
	"testing"
)

func TestInitConfig(t *testing.T) {
	if err := InitConfig(); err != nil {
		log.Error(err)
		return
	}
	if Conf == nil{
		fmt.Println("Config初始化错误")
	}else {
		fmt.Println("Config初始化正常")
	}
}

func TestClientConfig_GetMysqlConfig(t *testing.T) {
	if err := InitConfig(); err != nil {
		log.Error(err)
		return
	}
	inst := Conf.GetMysqlConfig()
	fmt.Println(inst)
}