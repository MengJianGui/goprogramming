package config

import "github.com/astaxie/beego/config"

//MysqlInstance mysql instance.
type MysqlInstance struct {
	Host     string
	Port     string
	DBName   string
	UserName string
	Passwd   string
	Charset  string
}

// 通过beego/config包中的config模块来解析你各种格式的文件
type ClientConfig struct {
	Configer config.Configer
}
//type Configer config.Configer

var (
	Conf *ClientConfig = nil
)

// NewConfig read config from config.ini
func NewConfig() (*ClientConfig, error) {
	// “../”来表示上一级目录; “./”表示当前目录
	conf, err := config.NewConfig("ini", "../config/config.ini")
	if err != nil {
		return nil, err
	}

	return &ClientConfig{Configer: conf}, nil
}

//InitConfig init config
func InitConfig() error {
	conf, err := NewConfig()
	if err != nil {
		return err
	}
	Conf = conf
	return nil
}

// GetMySQLConfig get dbs mysql config
func (c *ClientConfig) GetMysqlConfig() *MysqlInstance {
	return &MysqlInstance{
		Host:     c.Configer.String("database::host"),
		Port:     c.Configer.String("database::port"),
		DBName:   c.Configer.String("database::name"),
		UserName: c.Configer.String("database::user"),
		Passwd:   c.Configer.String("database::pwd"),
		Charset:  c.Configer.String("database::charset"),
	}
}
