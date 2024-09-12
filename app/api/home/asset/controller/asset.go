package controller

import "github.com/gin-gonic/gin"

type AssetController struct {
}

func (ac AssetController) AssetDemo(c *gin.Context) {
	c.String(200, "AssetDemo ok")
}
