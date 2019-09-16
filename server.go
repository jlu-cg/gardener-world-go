package main

import (
	"fmt"

	"github.com/gardener/gardener-world-go/config"
	"github.com/gardener/gardener-world-go/routers"
	"github.com/gardener/gardener-world-go/service"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

func main() {

	config := config.New()
	config.LoadConfig()
	service.InitPool(config)
	service.InitCouchDb(config)

	fmt.Printf("%d\n", config.ServerConfig.Port)
	fmt.Printf("%s", config.PgConfig.URL)

	app := iris.New()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	routers.InitDocRouter(app, crs)

	app.Run(iris.Addr(":38080"))
}
