package routers

import (
	"net/http"

	"../databaseConfig"
	"../model"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := databaseConfig.DbDetails()
		defer session.Close()

		username, password, ok := c.Request.BasicAuth()
		auth := model.CheckUserNameAndPassword(username, password, session)
		if !ok || auth != true {
			c.Abort()
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Writer.Write([]byte("Unauthorized"))
			return
		}
	}
}

func Create(c *gin.Context) {
	//var data []model.UserDetails
	session := databaseConfig.DbDetails()
	defer session.Close()
	deserilizesChannel := make(chan []model.UserDetails)
	serilizeChannel := make(chan []model.UserDetails)
	go Deserializations(c, deserilizesChannel)
	//c.BindJSON(&data)
	data, err := model.InsertData(<-deserilizesChannel, serilizeChannel, session)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

func Read(c *gin.Context) {
	session := databaseConfig.DbDetails()
	defer session.Close()
	id := c.Params.ByName("id")
	data, err := model.GetUser(id, session)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

func Update(c *gin.Context) {
	session := databaseConfig.DbDetails()
	defer session.Close()
	id := c.Params.ByName("id")
	desUpdateChannel := make(chan model.UserDetails)
	go DeserializationUpdate(c, desUpdateChannel)
	data, err := model.UpdateUser(id, <-desUpdateChannel, session)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

func Delete(c *gin.Context) {
	session := databaseConfig.DbDetails()
	defer session.Close()
	id := c.Params.ByName("id")
	err := model.DeleteUser(id, session)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}

func ReadAll(c *gin.Context) {
	session := databaseConfig.DbDetails()
	defer session.Close()
	data, err := model.GetUserAll(session)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, data)
	}
}
