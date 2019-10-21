package routers

import (
	"../model"
	"github.com/gin-gonic/gin"
)

func Deserializations(c *gin.Context, deserilizesChannel chan []model.UserDetails) {
	var data []model.UserDetails
	c.BindJSON(&data)
	deserilizesChannel <- data

}
func DeserializationUpdate(c *gin.Context, desUpdateChannel chan model.UserDetails) {
	var data model.UserDetails
	c.BindJSON(&data)
	desUpdateChannel <- data

}
