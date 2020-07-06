package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"goprogramming/src/mysql/config"
)

/**
测试MySQL数据库连接
*/

var (
	// sql.DB表示操作数据库的抽象访问接口,但不是所谓的数据库连接对象.
	// sql.DB的设计就是用来作为长连接使用的。不要频繁Open, Close。
	//比较好的做法是，为每个不同的datastore建一个DB对象，保持这些对象Open。
	//如果需要短连接，那么把DB作为参数传入function，而不要在function中Open, Close。
	DB *sql.DB
)

// NewMysqlConn open new connect to db
func NewMysqlConn(inst *config.MysqlInstance) (*sql.DB, error) {
	dbinfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		inst.UserName, inst.Passwd, inst.Host, inst.Port, inst.DBName, inst.Charset)
	//sql.Open并不会立即建立一个数据库的网络连接, 也不会对数据库链接参数的合法性做检验,
	//它仅仅是初始化一个sql.DB对象. 当真正进行第一次数据库查询操作时, 此时才会真正建立网络连接;
	db, err := sql.Open("mysql", dbinfo)
	fmt.Printf("data source name: %s", dbinfo)
	fmt.Println()
	if err != nil {
		return nil, err
	}
	// 通过Ping()函数来判断是否连接上了数据库
	if err := db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(0)
	return db, nil
}

// InitDbConn初始化一个DB长连接
func InitDbConn() error {
	dbinst := config.Conf.GetMysqlConfig()
	db, err := NewMysqlConn(dbinst)
	if err != nil {
		return err
	}

	DB = db
	return nil
}
