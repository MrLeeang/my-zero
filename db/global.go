package db

import (
	"context"
	"encoding/json"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var GlobalSession *gorm.DB

// 初始化数据库连接池
func createSession(msn string) (*gorm.DB, error) {

	var err error
	GlobalSession, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                      msn,
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名禁用复数
		},
		Logger: LogxLogger{
			level:                     logger.Info,
			IgnoreRecordNotFoundError: false,           // 忽略not found
			SlowThreshold:             1 * time.Second, // 慢sql
		},
	})

	if err != nil {
		return nil, err
	}

	sqlDb, err := GlobalSession.DB()
	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxIdleConns(20)   //设置最大空闲连接数
	sqlDb.SetMaxOpenConns(100)  //设置最大打开连接数
	sqlDb.SetConnMaxLifetime(0) // 设置连接的最大生命周期

	GlobalSession.AutoMigrate(new(SysUser))

	return GlobalSession, nil
}

func InitializeDatabase(msn string) *gorm.DB {

	if GlobalSession != nil {
		// 关闭已经打开的数据库连接
		sqlDB, err := GlobalSession.DB()
		if err == nil {
			sqlDB.Close()
		}
	}

	session, err := createSession(msn)

	if err != nil {
		panic(err)
	}

	return session
}

func ToJson(model interface{}) string {

	resByte, _ := json.Marshal(model)

	return string(resByte)
}

func Add(ctx context.Context, model interface{}) error {
	return GlobalSession.WithContext(ctx).Create(model).Error
}

func Delete(ctx context.Context, model interface{}, where ...interface{}) error {
	return GlobalSession.WithContext(ctx).Delete(model, where...).Error
}

func Unscoped(ctx context.Context, model interface{}, where ...interface{}) error {
	return GlobalSession.WithContext(ctx).Unscoped().Delete(model, where...).Error
}

func Save(ctx context.Context, model interface{}) error {
	return GlobalSession.WithContext(ctx).Save(model).Error
}

func First(ctx context.Context, model interface{}, where ...interface{}) error {
	return GlobalSession.WithContext(ctx).First(model, where...).Error
}

func List(ctx context.Context, model interface{}, where ...interface{}) error {
	return GlobalSession.WithContext(ctx).Find(model, where...).Error
}

func ListOrderByID(ctx context.Context, model interface{}, where ...interface{}) error {
	return GlobalSession.WithContext(ctx).Order("id desc").Find(model, where...).Error
}
