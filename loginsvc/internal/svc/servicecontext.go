package svc

import (
	"github.com/MrLeeang/my-zero/db"
	"github.com/MrLeeang/my-zero/loginsvc/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 加载全局数据库，go-zero自带的sqlx，不好用，不如gorm习惯
	// database.LoadDatabase(c.Mysql.Dsn)
	// 加载全局数据库
	session := db.InitializeDatabase(c.Mysql.Dsn)

	return &ServiceContext{
		Config: c,
		DB:     session,
	}
}
