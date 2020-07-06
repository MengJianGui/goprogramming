package database

import (
	"fmt"
	"github.com/zssky/log"
)

/**
数据库曾查改删操作insert、select、update、delete
*/

type UserInfo struct {
	Name string
	Age  int
}

// 插入数据，而如果通过预编译(prepared statement)插入数据的话一般适用于批量操作。
func InsertData(info *UserInfo) error {
	sql := `INSERT INTO ssm_hrm.user (name, age) VALUES (?, ?)`
	_, err := DB.Exec(sql, info.Name, info.Age)
	if err != nil {
		log.Errorf("InsertData{%s:%s} error{%v}", sql, info, err)
		return err
	}
	log.Infof("InsertData{%s:%s} success!", sql, info)
	return nil
}

// 查询数据
func SelectData() error {
	sql := `SELECT name, age FROM ssm_hrm.user WHERE age < ?`
	rows, err := DB.Query(sql, 35)
	defer rows.Close()
	// defer DB.Close()
	if err != nil {
		log.Errorf("SelectData{%s:35} error{%v}", sql, err)
		return err
	}
	log.Infof("SelectData{%s:35} success!", sql)
	var result []*UserInfo
	for rows.Next() {
		//columns, _ := rows.Columns()
		info := &UserInfo{}
		// 传入的一定是地址值
		err := rows.Scan(&info.Name, &info.Age) // mysql的int型转换为go的int64类型
		if err != nil {
			log.Errorf("rows.Scan error{%v}", err)
			return err
		}

		result = append(result, info)
	}
	fmt.Println(result)
	return nil
}

// 修改数据update
func UpdataData() error {
	sql := `UPDATE ssm_hrm.user SET age = ? WHERE name = ?`
	_, err := DB.Exec(sql, 18, "mengjiangui")
	if err != nil {
		log.Errorf("Update Data{%s} error{%v}", sql, err)
		return err
	}
	log.Infof("Update Data{%s} success!", sql)
	return nil
}

// 实现Stringer接口的String()方法，输出结构体指针变量
func (info *UserInfo) String() string {
	return fmt.Sprintf("UserInfo{Name:%s, Age:%d}", info.Name, info.Age)
}
