/**
  @author:panliang
  @data:2022/5/27
  @note
**/
package bootstrap

import (
	"github.com/gin-gonic/gin"
	router2 "im-services/app/router"
	"im-services/app/service/client"
	"im-services/app/service/queue/nsq_queue"
	"im-services/config"
	"im-services/pkg/coroutine_poll"
	"im-services/pkg/logger"
	"im-services/pkg/model"
	"im-services/pkg/nsq"
	"im-services/pkg/redis"
	"net/http"
	_ "net/http/pprof"
)

// 启动服务方法
func Start() {

	r := gin.Default()

	go client.ImManager.Start()

	setRoute(r)

	gin.SetMode(config.Conf.Server.Mode)

	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()
	_ = r.Run(config.Conf.Server.Listen)
}

// 启动第三方服务
func LoadConfiguration() {

	setUpLogger()

	model.InitDb()

	redis.InitClient()

	coroutine_poll.ConnectPool()

	_ = nsq.InitNewProducerPoll()
	go nsq_queue.ConsumersInit()

}

// 初始化日志方法
func setUpLogger() {
	logger.InitLogger(
		config.Conf.Log.FileName,
		config.Conf.Log.MaxSize,
		config.Conf.Log.MaxBackup,
		config.Conf.Log.MaxAge,
		config.Conf.Log.Compress,
		config.Conf.Log.Type,
		config.Conf.Log.Level,
	)
}

// 注册路由方法
func setRoute(r *gin.Engine) {
	router2.RegisterApiRoutes(r)
	router2.RegisterWsRouters(r)
}