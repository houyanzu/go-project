package controller

import "github.com/gin-gonic/gin"

type UserController struct {
}

func (uc UserController) Demo(c *gin.Context) {
	c.String(200, "demo ok")
}

func (uc UserController) Demo2(c *gin.Context) {
	c.String(200, "demo2 ok")
}
