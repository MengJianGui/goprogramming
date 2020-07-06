/**
beego操作数据库
 */

package main

import (
	"fmt"
	"strings"
	"time"

	// 导入orm包
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/zssky/log"
	// 导入mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

// Resource
type Resource struct {
	Host       string `json:"host",orm:"host"`
	User       string `json:"user",orm:"user"`
	Pass       string `json:"pass",orm:"pass"`
	Category   string `json:"category",orm:"category"`
	Status     string `json:"status",orm:"status"`
	Message    string `json:"message",orm:"message"`
	CreateTime string `json:"createTime",orm:"create_time"`
	UpdateTime string `json:"updateTime",orm:"update_time"`
}

// 根据恢复机器人ip判断当前机器是否正在进行恢复
func GetRecoverStatus(host string) (*Resource, error) {
	query := `select create_time from t_backup_resources where host=?;`

	log.Infof("query %s", query)
	var rs  *Resource
	o := orm.NewOrm()
	// beego 没有查到数据的话会报错。
	if err := o.Raw(query, host).QueryRow(&rs); err != nil {
		log.Errorf("no recover resource{%v}", err)
		return rs, err
	}

	return rs, nil
}


func MM(p interface{}) {
	mm := p.(*Resource)
	fmt.Printf("p interface{%+v}", mm)
}

// 通过init函数配置mysql数据库连接信息
func init() {
	// 这里注册一个default默认数据库，数据库驱动是mysql.
	// 第三个参数是数据库dsn, 配置数据库的账号密码，数据库名等参数
	//  dsn参数说明：
	//      username    - mysql账号
	//      password    - mysql密码
	//      db_name     - 数据库名
	//      127.0.0.1:3306 - 数据库的地址和端口
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/web?charset=utf8&parseTime=true&loc=Local")

	// 打开调试模式，开发的时候方便查看orm生成什么样子的sql语句
	orm.Debug = true

	rs, err := GetRecoverStatus("11.38.125.251")
	switch err {
	case nil:
		fmt.Println("err is nil")
	default:
		if strings.Contains(err.Error(), "no row found") {
			fmt.Println(err.Error())
		}
		return
	}
	preTime, err := time.ParseInLocation("2006-01-02 15:04:05 +0800 CST", rs.CreateTime, time.Local)
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Println(preTime.Unix())
	fmt.Printf("time.Now().Unix(): %v", time.Now().Unix())
	fmt.Println()
}

func main() {
	beego.Run()
}
