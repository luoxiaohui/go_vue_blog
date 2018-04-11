var domain = process.env.API_DOMAIN
var port = process.env.API_PORT

global.articleList = domain + ":" + port + "/articleList"
global.articleDetail = domain + ":" + port + "/articleDetail"