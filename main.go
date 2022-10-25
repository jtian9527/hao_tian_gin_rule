
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"haotian_rule/config"
	"haotian_rule/routers"
	"haotian_rule/serviceGrpc"
	"haotian_rule/utils"
	"net"
)

func main(){
	go useGrpc()
	useGin()
}

func useGin() {
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

func useGrpc() {
	const port = ":9091" // 服务器端口
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	serviceGrpc.RegisterDemoServiceServer(s, &serviceGrpc.MyDemoServer{})
	//reflection.Register(s)

	defer func() {
		s.Stop()
		listen.Close()
	}()
	err = s.Serve(listen)
	if err != nil {
		fmt.Printf("failed to server: %v", err)
		return
	}
}