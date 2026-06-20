package handlers
<<<<<<< main
=======

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Sukanthan06/code-execution-engine/executor/sandbox"

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

	output, err := sandbox.RunPython(req.Code)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"output": output,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"output": output,
	})
}
>>>>>>> local
