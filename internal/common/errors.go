package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var ForbiddenError = errors.New("not allowed")

func HandleError(ctx *gin.Context, err error) {
	if errors.Is(err, ForbiddenError) {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
