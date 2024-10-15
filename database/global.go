package database

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	Conn sqlx.SqlConn
)

func LoadDatabase(dsn string) sqlx.SqlConn {
	Conn = sqlx.NewSqlConn("mysql", dsn)
	return Conn
}
