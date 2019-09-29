package load

import (
	"github.com/silenceper/wechat/cache"
	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/config"
)

var Cache *cache.Redis

func Redis() {
	opts := &cache.RedisOpts{
		Host:        config.RedisHost,
		Password:    config.RedisPassword,
		Database:    config.RedisDatabase,
		MaxIdle:     config.MaxIdle,
		MaxActive:   config.MaxActive,
		IdleTimeout: config.IdleTimeout,
	}
	Cache = cache.NewRedis(opts)
}
