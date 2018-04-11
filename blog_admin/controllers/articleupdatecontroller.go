package controllers

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt" 
	"strings"
)

type ArticleUpdateController struct {
	beego.Controller
}

type ArticleUpdate struct{
	ArticleId string `json:"articleId"` //首字母大写表示public，小写表示private，添加tag，使其输出时变成小写
	Title string `json:"title"`
	Date string `json:"date"`
	Content string `json:"content"`
	Gist string `json:"gist"`
	Labels string `json:"labels"`
}

func (this *ArticleUpdateController) Post() {
	isLogin := this.GetSession("isLogin")
	fmt.Println("准备打印isLogin....")
	fmt.Println(isLogin)

	result := make(map[string]interface{})
	if isLogin == nil{
		result["code"] = "001"
		result["msg"] = "fail"
		result["data"] = "登录失效，请重新登录"
		
	}else{
		var ob ArticleUpdate
		json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	
		var isUpdateSuccess bool = updateArticle(ob.ArticleId,ob.Title,ob.Date,ob.Content,ob.Gist,ob.Labels)
		result["msg"] = "success"
		if isUpdateSuccess{
			result["code"] = "000"
			result["data"] = "文章更新成功"
		}else{
			result["code"] = "002"			
			result["data"] = "文章更新失败"
		}
	}
	bytes, err := json.Marshal(result)
	if err != nil{
		fmt.Println(err)
	}
	this.Data["json"] = string(bytes)
	this.ServeJSON()
}

//查询文章列表
func updateArticle(articleId string,title string,date string, content string,gist string,lables string) bool{
	var dbhost string = beego.AppConfig.String("dbhost")
	var dbport string = beego.AppConfig.String("dbport")
	var dbuser string = beego.AppConfig.String("dbuser")
	var dbpassword string = beego.AppConfig.String("dbpassword")
	var dbname string = beego.AppConfig.String("dbname")
	var dbcharset string = beego.AppConfig.String("dbcharset")

	var isUpdateSuccess bool = false
	db, err := sql.Open("mysql", dbuser+":"+dbpassword+"@tcp("+dbhost+":"+dbport+")/"+dbname+"?"+dbcharset)
	if(err != nil){
		fmt.Println(err)
	}

	content = strings.Replace(content,`"` , `\"`, -1)
	content = strings.Replace(content,`'` , `\'`, -1)
	content = strings.Replace(content,`\` , `\\\`, -1)
	var sql string = "update article set title = " +  "\""+title+"\", content = " + "\""+content+"\", date = " + "\""+date+"\", gist = " + "\""+gist+"\", labels = " + "\""+ lables + "\" where articleId = \"" + articleId + "\""
	fmt.Println(sql)
	rows, err := db.Query(sql)
	if(err != nil){ 
		fmt.Println(err)
	}
	
	for rows.Next() {
		
		isUpdateSuccess = true
	}
	rows.Close()
	db.Close()
	return isUpdateSuccess
}