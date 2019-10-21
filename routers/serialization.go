package routers

import (
	"../model"
	"github.com/gin-gonic/gin"
)

func Serializations(c *gin.Context, serilizeChannel chan []model.UserDetails) {
	var data []model.UserDetails
	c.BindJSON(&data)
	serilizeChannel <- data

}
