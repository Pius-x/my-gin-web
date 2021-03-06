package initialize

import (
	"github.com/my-gin-web/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/global"
	"github.com/my-gin-web/router"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()
	routerGroup := router.GroupApp

	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面4行注释
	// Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/static", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	Router.LoadHTMLGlob("resource/view/*")
	Router.StaticFS(global.Config.Local.Path, http.Dir(global.Config.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	global.ZapLog.Info("use middleware logger")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	//Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	global.ZapLog.Info("use middleware cors")

	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}

	PublicGroup.Use(middleware.InterceptPublic())
	{
		routerGroup.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}

	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.OperationRecord()).Use(middleware.InterceptPrivate())
	{
		routerGroup.InitUserRouter(PrivateGroup)         // 注册用户路由
		routerGroup.InitAuthorityRouter(PrivateGroup)    // 注册角色路由
		routerGroup.InitOperationLogRouter(PrivateGroup) // 操作记录

		//exampleRouter.InitExcelRouter(PrivateGroup)                 // 表格导入导出
		//exampleRouter.InitCustomerRouter(PrivateGroup)              // 客户路由
		//exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
	}

	InstallPlugin(PublicGroup, PrivateGroup) // 安装插件

	global.ZapLog.Info("router register success")
	return Router
}
