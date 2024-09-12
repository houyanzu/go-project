package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//var configName string
	var port string = "8888"
	//flag.StringVar(&port, "port", "8022", "port")
	//flag.StringVar(&configName, "config", "", "Config json file name")
	//flag.Parse()
	//err := myconfig.ParseConfig(configName)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.InitMysql()
	//if err != nil {
	//	panic(err)
	//}
	//err = cache.InitCache()
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = boxlog.Init("home_api", true)
	//if err != nil {
	//	panic(err)
	//}

	router := gin.Default()
	Register(router)
	err := router.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
