package handlers

import (
	"net/http"

	"github.com/Sukanthan06/code-execution-engine/executor/sandbox"
	"github.com/Sukanthan06/code-execution-engine/languages"
	"github.com/Sukanthan06/code-execution-engine/models"
	"github.com/gin-gonic/gin"
)

func ExecuteCode(c *gin.Context) {
	var req models.ExecuteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	lang, exists := languages.SupportedLanguages[req.Language]

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unsupported language",
		})
		return
	}

	output, err := sandbox.RunCode(req.Code, lang.Extension, lang.Interpreter, lang.Compiler)

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
