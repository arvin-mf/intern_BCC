package response

import (
	"intern_BCC/model"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, httpCode int, msg string, data interface{}, pagin *model.PaginParam) {
	c.JSON(httpCode, map[string]interface{}{
		"status":     "success",
		"message":    msg,
		"data":       data,
		"pagination": pagin,
	})
}

func FailOrError(c *gin.Context, httpCode int, msg string, err error) {
	c.JSON(httpCode, gin.H{
		"status":  "fail",
		"message": msg,
		"data": gin.H{
			"error": err.Error(),
		},
	})
}
