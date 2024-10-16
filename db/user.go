package db

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	Uuid               string `db:"uuid"`
	Name               string `db:"name"`
	Gender             string `db:"gender"`
	LoginUser          string `db:"login_user"`
	LoginPass          string `db:"login_pass"`
	Email              string `db:"email"`
	PhoneNum           string `db:"phone_num"`
	IdNum              string `db:"id_num"`
	RoleUuid           string `db:"role_uuid"`
	TeamUuid           string `db:"team_uuid"`
	Enable             int64  `db:"enable"`
	Picture            string `db:"picture"`
	IsOnline           int64  `db:"is_online"`
	GroupUuid          string `db:"group_uuid"`
	ClassUuid          string `db:"class_uuid"`
	IsStudying         int64  `db:"is_studying"`
	LatestLoginErrTime int64  `db:"latest_login_err_time"`
}
