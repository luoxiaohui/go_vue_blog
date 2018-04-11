package routers

import (
	"go_vue_blog/blog_admin/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	// 这段代码放在router.go文件的init()的开头
    beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: 	  []string{"http://"+beego.AppConfig.String("front_end_domain")+":"+beego.AppConfig.String("front_end_port")},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
        AllowCredentials: true,
	}))
	
	// beego.Router("/admin/signup",&controllers.MainController{})
	beego.Router("/admin/signin",&controllers.LoginController{})
	beego.Router("/admin/saveArticle",&controllers.ArticleSaveController{})
	beego.Router("/admin/updateArticle",&controllers.ArticleUpdateController{})
	beego.Router("/admin/deleteArticle",&controllers.ArticleDeleteController{})
	beego.Router("/admin/articleDetail",&controllers.ArticleDetailController{})
	beego.Router("/admin/articleList",&controllers.ArticleListController{})
	
	
}
