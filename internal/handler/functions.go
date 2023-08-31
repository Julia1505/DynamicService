package handler

import "github.com/gin-gonic/gin"

func GetSegmentFromContext(ctx *gin.Context) (string, error) {
	var req RequestSegment
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return "", err
	}
	return string(req), nil
}
