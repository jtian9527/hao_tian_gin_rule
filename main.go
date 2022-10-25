
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"haotian_rule/config"
	"haotian_rule/routers"
	"haotian_rule/utils"
)


func main() {
	conf, err := config.ParseConfig("./config/config.json")
	if err != nil {
		panic("读取配置文件失败，" + err.Error())
	}
	fmt.Printf("conf:%#v\n", conf)
	utils.InitUcache(conf.RedisConfig)
	engine := gin.Default()
	routers.RegisterRouter(engine)
	engine.Run(":8012")
}
