var domain = process.env.API_DOMAIN
var port = process.env.API_PORT

global.signin = domain + ":" + port + "/admin/signin"
global.saveArticle = domain + ":" + port + "/admin/saveArticle"
global.updateArticle = domain + ":" + port + "/admin/updateArticle"
global.deleteArticle = domain + ":" + port + "/admin/deleteArticle"
global.articleDetail = domain + ":" + port + "/admin/articleDetail"
global.articleList = domain + ":" + port + "/admin/articleList"