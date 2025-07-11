package resputil

import (
	"github.com/gin-gonic/gin"
)

func WriteResponse(ctx *gin.Context, err error, data interface{}) {
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, data)
}
