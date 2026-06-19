package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Sukanthan06/code-execution-engine/models"
)

func ExecuteCode(c *gin.Context) {
	var req models.ExecuteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"language": req.Language,
		"code":     req.Code,
	})
}
