package gredis

import (
	"encoding/json"
	"github.com/EDDYCJY/go-gin-demo/pkg/setting"
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisConn *redis.Pool // 初始化 redis池

func Setup() error {
	// redis池化 -> 采取地址&给当前结构体赋值
	RedisConn = &redis.Pool{
		MaxIdle:     setting.RedisSetting.MaxIdle,     // 连接数
		MaxActive:   setting.RedisSetting.MaxActive,   // 活跃数
		IdleTimeout: setting.RedisSetting.IdleTimeout, // 超时
		Dial: func() (redis.Conn, error) { // 拨号或实例化redis链接
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				// 如果密码存在 -> 需要加上密码 -> c.Do 发送指令
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		// 心跳测试 -> c.Do 发送指令 -> ping 是否能通 -> 失败返回 err
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("ping")
			return err
		},
	}
	return nil
}

// Set 设置值、设置超时
func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("Expire", key, time)
	if err != nil {
		return err
	}

	return nil
}

// Exists 检查key是否存在
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()
	exist, err := redis.Bool(conn.Do("EXSITS", key))
	if err != nil {
		return false
	}
	return exist
}
