package controllers

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt" 
)

type ArticleDeleteController struct {
	beego.Controller
}

type ArticleDetele struct{
	ArticleId string `json:"articleId"` //首字母大写表示public，小写表示private，添加tag，使其输出时变成小写
}


func (this *ArticleDeleteController) Post() {
	isLogin := this.GetSession("isLogin")
	fmt.Println("准备打印isLogin....")
	fmt.Println(isLogin)

	result := make(map[string]interface{})
	if isLogin == nil{
		result["code"] = "001"
		result["msg"] = "fail"
		result["data"] = "登录失效，请重新登录"
		
	}else{
		var ob ArticleDetele
		json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
		
		var isDeleteSuccess bool = deleteArticle(ob.ArticleId)
		if isDeleteSuccess{
			result["code"] = "000"
			result["msg"] = "success"
			result["data"] = "文章删除成功"
		}else{

			result["code"] = "002"
			result["msg"] = "faile"
			result["data"] = "文章删除失败"
		}
		
	}
	bytes, err := json.Marshal(result)
	if err != nil{
		fmt.Println(err)
	}
	this.Data["json"] = string(bytes)
	this.ServeJSON()
}

// 删除文章
func deleteArticle(articleId string) bool{
	var dbhost string = beego.AppConfig.String("dbhost")
	var dbport string = beego.AppConfig.String("dbport")
	var dbuser string = beego.AppConfig.String("dbuser")
	var dbpassword string = beego.AppConfig.String("dbpassword")
	var dbname string = beego.AppConfig.String("dbname")
	var dbcharset string = beego.AppConfig.String("dbcharset")

	var isDeleteSuccess bool = false
	db, err := sql.Open("mysql", dbuser+":"+dbpassword+"@tcp("+dbhost+":"+dbport+")/"+dbname+"?"+dbcharset)
	if(err != nil){
		fmt.Println(err)
	}

	var sql string = "delete FROM article where articleId = \"" + articleId + "\""
	fmt.Println(sql)
	rows, err := db.Query(sql)
	if(err != nil){
		fmt.Println(err)
	}
	
	for rows.Next() {
		
        isDeleteSuccess = true
	}

	rows.Close()
	db.Close()
	return isDeleteSuccess
}