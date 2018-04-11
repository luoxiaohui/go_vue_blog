package controllers

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt" 
	"strings"
	"crypto/md5"
	"encoding/hex"
)

type ArticleSaveController struct {
	beego.Controller
}

type ArticleSave struct{
	ArticleId string `json:"articleId"` //首字母大写表示public，小写表示private，添加tag，使其输出时变成小写
	Title string `json:"title"`
	Date string `json:"date"`
	Content string `json:"content"`
	Gist string `json:"gist"`
	Labels []string `json:"labels"`
}

func (this *ArticleSaveController) Post() {
	isLogin := this.GetSession("isLogin")
	fmt.Println("保存文章接口...")
	fmt.Println(isLogin)

	result := make(map[string]interface{})
	if isLogin == nil{
		result["code"] = "001"
		result["msg"] = "fail"
		result["data"] = "登录失效，请重新登录"
		
	}else{
		var ob ArticleSave
		json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	
		fmt.Println("fuck you...")
		fmt.Println(ob.Labels)
		var isSaveSuccess bool = saveArticle(ob.Title,ob.Date,ob.Content,ob.Gist,ob.Labels)
		result["msg"] = "success"
		if isSaveSuccess{
			result["code"] = "000"
			result["data"] = "文章保存成功"
		}else{
			result["code"] = "002"			
			result["data"] = "文章保存失败"
		}
	}
	bytes, err := json.Marshal(result)
	if err != nil{
		fmt.Println(err)
	}
	this.Data["json"] = string(bytes)
	this.ServeJSON()
}

//保存文章
func saveArticle(title string,date string, content string,gist string,labels []string) bool{
	var dbhost string = beego.AppConfig.String("dbhost")
	var dbport string = beego.AppConfig.String("dbport")
	var dbuser string = beego.AppConfig.String("dbuser")
	var dbpassword string = beego.AppConfig.String("dbpassword")
	var dbname string = beego.AppConfig.String("dbname")
	var dbcharset string = beego.AppConfig.String("dbcharset")

	fmt.Println(len(labels))
	var labelsString string = ""
	for i := 0; i < len(labels); i++{
		labelsString += "," + labels[i]
	}
	rs := []rune(labelsString)
	labelsString = string(rs[1:len(labelsString)])
	
	fmt.Println(22222)
	var isSaveSuccess bool = false
	db, err := sql.Open("mysql", dbuser+":"+dbpassword+"@tcp("+dbhost+":"+dbport+")/"+dbname+"?"+dbcharset)
	if(err != nil){
		fmt.Println(err)
	}

	h := md5.New()
	h.Write([]byte(date))
	if err != nil{

	}
	fmt.Println(3333)
	// articleId用日期md5加密后的数据
	var articleId string = hex.EncodeToString(h.Sum(nil))
	content = strings.Replace(content,`"` , `\"`, -1)
	content = strings.Replace(content,`'` , `\'`, -1)
	content = strings.Replace(content,`\` , `\\\`, -1)
	fmt.Println(4444)
	var sql string = "insert into article(articleId,title,date,content,gist,labels) values(" + "\""+articleId+"\",\""+title+"\",\""+date+"\",\""+content+"\",\""+gist+"\",\""+labelsString+"\""+ ")"
	fmt.Println(sql)
	rows, err := db.Query(sql)
	if(err != nil){ 
		fmt.Println(err)
	}
	
	for rows.Next() {
		
		isSaveSuccess = true
	}
	rows.Close()
	db.Close()
	return isSaveSuccess
}