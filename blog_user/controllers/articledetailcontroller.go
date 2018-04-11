package controllers

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt" 
	"strings"
)

type ArticleDetailController struct {
	beego.Controller
}

type ArticleDetail struct{
	ArticleId string `json:"articleId"` //首字母大写表示public，小写表示private，添加tag，使其输出时变成小写
	Title string `json:"title"`
	Date string `json:"date"`
	Content string `json:"content"`
	Gist string `json:"gist"`
	Labels string `json:"labels"`
}

//传到客户端的参数中，labels需要string数组
type ArticleDetailResult struct{
	ArticleId string `json:"articleId"` //首字母大写表示public，小写表示private，添加tag，使其输出时变成小写
	Title string `json:"title"`
	Date string `json:"date"`
	Content string `json:"content"`
	Gist string `json:"gist"`
	Labels []string `json:"labels"`
}

func (this *ArticleDetailController) Post() {
	result := make(map[string]interface{})
	
	var ob ArticleDetail
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	
	var articleDetail ArticleDetailResult = queryArticleDetail(ob.ArticleId)
	result["code"] = "000"
	result["msg"] = "success"
	result["data"] = articleDetail
	bytes, err := json.Marshal(result)
	if err != nil{
		fmt.Println(err)
	}
	this.Data["json"] = string(bytes)
	this.ServeJSON()
}

//查询文章详情
func queryArticleDetail(articleId string) ArticleDetailResult{
	var dbhost string = beego.AppConfig.String("dbhost")
	var dbport string = beego.AppConfig.String("dbport")
	var dbuser string = beego.AppConfig.String("dbuser")
	var dbpassword string = beego.AppConfig.String("dbpassword")
	var dbname string = beego.AppConfig.String("dbname")
	var dbcharset string = beego.AppConfig.String("dbcharset")

	var articleResult ArticleDetailResult
	var article ArticleDetail
	db, err := sql.Open("mysql", dbuser+":"+dbpassword+"@tcp("+dbhost+":"+dbport+")/"+dbname+"?"+dbcharset)
	if(err != nil){
		return articleResult
	}

	var sql string = "SELECT articleId,title,content,date,gist,labels FROM article where articleId = \"" + articleId + "\""
	fmt.Println(sql)
	rows, err := db.Query(sql)
	if(err != nil){
		return articleResult
	}
	
	for rows.Next() {
		
        rows.Columns()
		err = rows.Scan(&article.ArticleId, &article.Title,&article.Content, &article.Date, &article.Gist, &article.Labels)
		if(err != nil){
			return articleResult
		}

	}
	var content string = article.Content
	content = strings.Replace(content,`\"` , `"`, -1)
	content = strings.Replace(content,`\'` , `'`, -1)
	content = strings.Replace(content,`\\\` , `\`, -1)
	article.Content = content

	//解决乱码
	var pureLabelString string = strings.Replace(article.Labels,"\u0000","",-1)
	var labels[] string = strings.Split(pureLabelString,",")
	articleResult.Labels = labels
	articleResult.ArticleId = article.ArticleId
	articleResult.Content = article.Content
	articleResult.Date = article.Date
	articleResult.Gist = article.Gist
	articleResult.Title = article.Title

	rows.Close()
	db.Close()
	return articleResult
}