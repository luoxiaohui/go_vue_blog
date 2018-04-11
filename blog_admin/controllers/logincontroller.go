package controllers

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
)

type LoginController struct{
	beego.Controller
}

type LoginModel struct{
	Name string `json:"name"`
	Password string `json:"password"`
}

func (this *LoginController) Post(){

	var ob LoginModel
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)

	var isLogin bool = queryLoginValid(ob.Name, ob.Password)
	result := make(map[string]interface{})
	if isLogin{
		this.SetSession("isLogin", true)
		result["code"] = "000"
		result["msg"] = "success"
		result["data"] = "登录成功"
		
	}else{
		result["code"] = "001"
		result["msg"] = "fail"
		result["data"] = "登录失败，请重新登录"
	}
	bytes, err := json.Marshal(result)
	if err != nil{
		fmt.Println(err)
	}

	this.Data["json"] = string(bytes)
	this.ServeJSON()
}

func queryLoginValid(userName string, password string) bool{
	var dbhost string = beego.AppConfig.String("dbhost")
	var dbport string = beego.AppConfig.String("dbport")
	var dbuser string = beego.AppConfig.String("dbuser")
	var dbpassword string = beego.AppConfig.String("dbpassword")
	var dbname string = beego.AppConfig.String("dbname")
	var dbcharset string = beego.AppConfig.String("dbcharset")

	var isLogin bool = false
	db, err := sql.Open("mysql", dbuser+":"+dbpassword+"@tcp("+dbhost+":"+dbport+")/"+dbname+"?"+dbcharset)
	if(err != nil){
		fmt.Println(err)
	}

	var sql string = "SELECT name,password FROM user where name = " + "\""+userName +"\""+ "and password = " + "\""+password+"\""
	fmt.Println(sql)
	rows, err := db.Query(sql)
	if(err != nil){
		fmt.Println(err)
	}

	for rows.Next() {

		isLogin = true
	}

	rows.Close()
	db.Close()
	return isLogin
}
