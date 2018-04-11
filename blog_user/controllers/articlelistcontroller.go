package controllers

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
	"strings"
)

type ArticleListController struct { 
	beego.Controller
}

type Article struct{
	ArticleId string `json:"articleId"` //首字母大写表示public，小写表示private，添加tag，使其输出时变成小写
	Title string `json:"title"`
	Date string `json:"date"`
	Gist string `json:"gist"`
	Labels string `json:"labels"` 
}

type ArticleResult struct{
	ArticleId string `json:"articleId"` //首字母大写表示public，小写表示private，添加tag，使其输出时变成小写
	Title string `json:"title"`
	Date string `json:"date"`
	Gist string `json:"gist"`
	Labels []string `json:"labels"` 
}

func (this *ArticleListController) Post() {

	result := make(map[string]interface{})
	
	result["code"] = "000"
	result["msg"] = "success"
	var articleList []ArticleResult = queryArticleList() 
	result["data"] = articleList

	bytes, err := json.Marshal(result)
	if err != nil{
		fmt.Println(err)
	}

	this.Data["json"] = string(bytes)
	this.ServeJSON()
}

//查询文章列表
func queryArticleList() []ArticleResult{
	var dbhost string = beego.AppConfig.String("dbhost")
	var dbport string = beego.AppConfig.String("dbport")
	var dbuser string = beego.AppConfig.String("dbuser")
	var dbpassword string = beego.AppConfig.String("dbpassword")
	var dbname string = beego.AppConfig.String("dbname")
	var dbcharset string = beego.AppConfig.String("dbcharset")
	
	db, err := sql.Open("mysql", dbuser+":"+dbpassword+"@tcp("+dbhost+":"+dbport+")/"+dbname+"?"+dbcharset)
	if(err != nil){
		return nil
	}

	var sql string = "SELECT articleId,title,date,gist,labels FROM article"
	fmt.Println(sql)
	rows, err := db.Query(sql)
	if(err != nil){
		return nil
	}

	var articleResults []ArticleResult
	for rows.Next() {
		
		var article Article
		var articleResult ArticleResult
        rows.Columns()
		err = rows.Scan(&article.ArticleId, &article.Title, &article.Date, &article.Gist, &article.Labels)
		if(err != nil){
			return nil
		}
		var pureLabelString string = strings.Replace(article.Labels,"\u0000","",-1)
		var labels[] string = strings.Split(pureLabelString,",")
		articleResult.Labels = labels
		articleResult.ArticleId = article.ArticleId
		articleResult.Date = article.Date
		articleResult.Gist = article.Gist
		articleResult.Title = article.Title

		articleResults = append(articleResults, articleResult)
	}
	rows.Close()
	db.Close()
	return articleResults
}
