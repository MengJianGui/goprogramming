/**
通过二进制包，往数据库多线程灌数据
 */

package main

import (
	"database/sql"
	"flag"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zssky/log"
)

var (
	Host   = flag.String("host", "127.0.0.1", "dump host")
	Port   = flag.Int("port", 5606, "dump port")
	User   = flag.String("user", "user", "user for dump")
	Pass   = flag.String("pass", "pass", "pass for dump user")
	Thread = flag.Int("thread", 4, "线程数")
	Loop   = flag.Int("loop", 10, "循环次数")
	DbName = flag.String("dbName", "test", "test database")
	Multi  = flag.Int("multi", 1, "multi insert")
)

//Db数据库连接池
var (
	DB   *sql.DB
	wait sync.WaitGroup
)

type Test struct {
	id   int64
	name string
	time string
}

func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		*User, *Pass, *Host, *Port, *DbName, "utf8")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}

func InsertTest(test Test) error {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return err
	}

	//准备sql语句
	insert := "INSERT INTO test1 (`name`,`column1`, `column2`,`column3`,`column4`,`column5`,`column6`,`column7`,`column8`,`column9`,`update_time`) VALUES (?,uuid(),uuid(),uuid(),uuid(),uuid(),uuid(),uuid(),uuid(),uuid(),now())"

	for i := 0; i < *Multi; i++ {
		//将参数传递到sql语句中并且执行
		_, err = tx.Exec(insert, fmt.Sprintf("%s-%d", test.name, i))
		if err != nil {
			log.Error(err)
			fmt.Println("Exec fail")
			return err
		}
	}
	//将事务提交
	return tx.Commit()
}

func LoopInsert() {

	//if err := DeleteData(); err != nil {
	//	log.Error(err)
	//}

	for i := 0; i < *Thread; i++ {
		wait.Add(1)
		go func(n int) {
			for j := 0; j < *Loop; j++ {
				test := Test{
					name: fmt.Sprintf("%s-%d-%d", "mengke", n, j),
				}
				if err := InsertTest(test); err != nil {
					log.Error(err)
					return
				}
			}
			wait.Done()
		}(i)
	}
	wait.Wait()
}

func DeleteData() error {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return err
	}

	//准备sql语句
	sql := "DELETE FROM test1 WHERE id > 1"

	_, err = tx.Exec(sql)
	//将事务提交
	return tx.Commit()
}

func main() {
	flag.Parse()
	log.Infof("host=%s,port=%d, user=%s, pass=%s, thread=%d, loop=%d, multi=%d, dbname=%s",*Host,*Port,*User,*Pass,*Thread,*Loop,*Multi,*DbName)
	log.Info("开始插入数据：",time.Now().Format("2006-01-02 15:04:05"))
	InitDB()
	LoopInsert()
	defer DB.Close()
	log.Info("结束插入数据：",time.Now().Format("2006-01-02 15:04:05"))
}
