package delivery

import (
	"7Zero4/config"
	"7Zero4/delivery/controller"
	"7Zero4/delivery/middleware"
	"7Zero4/manager"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type appServer struct {
	usecaseManager manager.UsecaseManager
	router         *gin.RouterGroup
	routerDev      *gin.RouterGroup
	engine         *gin.Engine
	host           string
}

// func (a *appServer) Router() *gin.Engine {
// 	return a.engine
// }

func Server() *appServer {
	router := gin.Default()
	appConfig := config.NewConfig()
	infra := manager.NewInfra(appConfig)
	repoManager := manager.NewRepositoryManager(infra)
	// repoManager.NotifRepo()
	usecaseManager := manager.NewUsecaseManager(repoManager)

	host := appConfig.Url
	// Add CORS middleware
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"Origin", "Date", "Content-Length", "Content-Type", "Content-Disposition", "Accept", "X-Requested-With", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Request-Method", "Access-Control-Request-Headers", "Authorization", "token"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	return &appServer{
		usecaseManager: usecaseManager,
		engine:         router,
		router:         router.Group("api/", middleware.NewAuthTokenMiddleware(usecaseManager.TokenUsecase()).RequiredToken()),

		routerDev: router.Group("activation/"),
		host:      host,
	}
}

// func initRouterConfiguration() *gin.Engine {
// 	gin.SetMode(gin.ReleaseMode)
// 	router := gin.Default()
// 	router.Use(corsConfiguration())
// 	return router
// }

// func corsConfiguration() gin.HandlerFunc {
// 	return cors.New(cors.Config{
// 		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
// 		AllowAllOrigins:  true,
// 		AllowHeaders:     []string{"Origin", "Date", "Content-Length", "Content-Type", "Content-Disposition", "Accept", "X-Requested-With", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Request-Method", "Access-Control-Request-Headers", "Authorization", "token"},
// 		AllowCredentials: false,
// 		MaxAge:           12 * time.Hour,
// 	})
// }

func (a *appServer) initControllers() {
	// buat daftarin controller ada disini
	// setiap controller, isinya harus ada isian dari usecaseManager
	controller.NewUserController(a.router, a.routerDev, a.usecaseManager.LoginUsecase(), a.usecaseManager.RegistUsecase())
}

func (a *appServer) Run() {
	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
