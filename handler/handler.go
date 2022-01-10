package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

func HandleError(ctx *gin.Context, msg string, status int) {

	ctx.JSON(status, gin.H{
		"code":    status,
		"ok":      false,
		"message": msg,
	})
	log.Println(msg)

}

func HandleSuccess(ctx *gin.Context, data string, status int) {

	ctx.JSON(status, gin.H{
		"code": status,
		"ok":   true,
		"data": data,
	})

}
