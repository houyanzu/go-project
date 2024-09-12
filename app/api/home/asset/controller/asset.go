package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/houyanzu/work-box/lib/output"
)

type AssetController struct {
}

func (ac AssetController) AssetDemo(c *gin.Context) {
	output.NewOutput(c, 1, nil).Out()
}
