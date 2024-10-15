package main

import (
	"github.com/gin-gonic/gin"
)

type vmReq struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

func main() {

	r := gin.Default()

	r.GET("/vm", func(ctx *gin.Context) {

		var params vmReq
		if err := ctx.ShouldBindQuery(&params); err != nil {
			ctx.JSON(200, gin.H{
				"data": params,
				"code": 1001,
			})
			return
		}
		ctx.JSON(200, gin.H{
			"data": params,
			"code": 0,
		})
	})

	r.Run(":8080")
}
