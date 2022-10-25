package controller

import (
	"github.com/gin-gonic/gin"
	"haotian_rule/config"
	"haotian_rule/utils"
	"haotian_rule/dao"
	"net/http"
	"time"
)


type UserController struct {
}
//新增用户
func (controller *UserController) Add(context *gin.Context) {
	name, exist := context.GetPostForm("name")
	if !exist || name == "" {
		context.JSON(http.StatusOK, gin.H{
			"msg": "请输入用户名:name",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}
//查询用户
func (controller *UserController) Get(context *gin.Context) {
	id := context.Query("id")
	currentTime := time.Now().Unix()
	UserKey :=  "haotian"
	err := utils.RedisClient.Set(UserKey, currentTime, 0).Err()
	if err != nil {
		panic(err)
	}
	demo := dao.GetDemoDao().GetDemoDaoName("haotian")
	val, err := utils.RedisClient.Get(UserKey).Result()
	if err != nil {
		panic(err)
	}
	context.JSON(http.StatusOK, gin.H{
		"id": id,
		"conf": config.GetConfig(),
		"redis_data":val,
		"db_data":	demo,

	})
}