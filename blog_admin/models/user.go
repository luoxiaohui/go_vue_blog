package models

import (
	// "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

// type User struct {
// 	name     string
// 	password string
// }

// func (m *User) TableName() string {
// 	return TableName("user")
// }


//查询数据
func query() {
    db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/my_blog?charset=utf8")
    checkErr(err)

    rows, err := db.Query("SELECT name,password FROM user")
    checkErr(err)

    for rows.Next() {
        var name string
        var password string

        rows.Columns()
        err = rows.Scan(&name, &password)
        checkErr(err)

        fmt.Println(name)
        fmt.Println(password)
    }
    db.Close()
}

func checkErr(err error){
	if err != nil {
        fmt.Println(err)
	}
}