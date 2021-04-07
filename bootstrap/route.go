package bootstrap

import (
	"awesomeProject/goweb/Handler"
	routes2 "awesomeProject/goweb/routes"
	"github.com/gorilla/mux"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	Handler.SetRoute(router)
	routes2.RegisterWebRoutes(router)
	return router
}
