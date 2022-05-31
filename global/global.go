package global

import (
	"github.com/jmoiron/sqlx"
	"github.com/my-gin-web/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"golang.org/x/sync/singleflight"

	"github.com/my-gin-web/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	Redis     *redis.Client
	ZapLog    *zap.Logger
	Config    config.Server
	_         *viper.Viper
	GormDbMap map[string]*gorm.DB
	DbMap     map[string]*sqlx.DB

	Timer              = timer.NewTimerTask()
	ConcurrencyControl = &singleflight.Group{}

	BlackCache local_cache.Cache
)
