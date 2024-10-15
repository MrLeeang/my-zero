package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/MrLeeang/my-zero/database"
)

func main() {
	// 加载全局数据库
	conn := database.LoadDatabase("root:51elab_mysql@tcp(192.168.2.235:3306)/merge_v1?charset=utf8mb4&parseTime=True&loc=Local")

	userModel := database.NewSysUserModel(conn)

	user, err := userModel.FindOne(context.TODO(), 1)

	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(user)

	fmt.Println(string(b))
}
