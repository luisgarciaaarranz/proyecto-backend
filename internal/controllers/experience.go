package controllers

import "github.com/gin-gonic/gin"

func GetExperienceHandler(c *gin.Context) {
	c.Data(200, "application/json", []byte(`hola`))
}
