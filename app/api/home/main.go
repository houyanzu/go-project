package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/houyanzu/work-box/database"
	"github.com/houyanzu/work-box/lib/boxlog"
	"github.com/houyanzu/work-box/tool/cache"
	_ "go-project/lib/outputmsg"
	"go-project/myconfig"
)

func main() {
	var configName string
	var port string
	flag.StringVar(&port, "port", "8022", "port")
	flag.StringVar(&configName, "config", "D:\\work\\gowork\\go-project\\myconfig\\config.json", "Config json file name")
	flag.Parse()
	err := myconfig.ParseConfig(configName)
	if err != nil {
		panic(err)
	}

	err = database.InitMysql()
	if err != nil {
		panic(err)
	}
	err = cache.InitCache()
	if err != nil {
		panic(err)
	}

	err = boxlog.Init("home_api", true)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	Register(router)
	err = router.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
