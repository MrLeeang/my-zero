package db

import (
	"context"
	"errors"
	"runtime/debug"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type LogxLogger struct {
	level                     logger.LogLevel
	IgnoreRecordNotFoundError bool
	SlowThreshold             time.Duration
}

func (l LogxLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.level = level
	return l
}

func (l LogxLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.level < logger.Info {
		return
	}
	logx.WithContext(ctx).Debugf(msg, data...)
}

func (l LogxLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.level < logger.Warn {
		return
	}
	logx.WithContext(ctx).Infof(msg, data...)
}

func (l LogxLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.level < logger.Error {
		return
	}
	logx.WithContext(ctx).Errorf(msg, data...)
}

func (l LogxLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	if l.level <= 0 {
		return
	}

	lg := logx.WithContext(ctx)
	sql, rows := fc()
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.level >= logger.Error && (!l.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		lg.Errorf("SQL: %s, Error: %s", sql, err.Error())
		lg.Error(string(debug.Stack()))
	case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.level >= logger.Warn:
		lg.Errorf("Warning: SQL: %s, Rows: %d, Duration: %s", sql, rows, time.Since(begin))
	default:
		lg.Infof("SQL: %s, Rows: %d, Duration: %s", sql, rows, time.Since(begin))
	}
}
