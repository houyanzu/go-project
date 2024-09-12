package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/houyanzu/work-box/lib/output"
	"github.com/houyanzu/work-box/tool/middleware"
)

type UserController struct {
}

func (uc UserController) Login(c *gin.Context) {
	token, _ := middleware.GetLoginToken(1, "", true)
	output.NewOutput(c, 0, gin.H{"token": token}).Out()
}

func (uc UserController) UserDemo2(c *gin.Context, userID uint) {
	output.NewOutput(c, 0, nil).DiyMsg("UserDemo2 ok" + fmt.Sprintf("%d", userID)).Out()
}
