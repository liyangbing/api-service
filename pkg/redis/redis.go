/**
  @author:panliang
  @data:2022/6/5
  @note
**/
package redis

import (
	"github.com/go-redis/redis"
	"im-services/config"
	"log"
	"time"
)

var RedisDB *redis.Client

// redis 连接
func InitClient() (err error) {

	RedisDB = redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         config.Conf.Redis.Host + ":" + config.Conf.Redis.Port,
		Password:     config.Conf.Redis.Password,
		DB:           config.Conf.Redis.DB,
		PoolSize:     config.Conf.Redis.Poll, //连接池 默认为4倍cpu数
		MinIdleConns: config.Conf.Redis.Conn, //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量
		DialTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		PoolTimeout:  5 * time.Second,
	})
	_, err = RedisDB.Ping().Result()

	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}