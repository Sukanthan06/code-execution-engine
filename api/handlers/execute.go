package handlers

import (
	"net/http"

	dockerrunner "github.com/Sukanthan06/code-execution-engine/executor/docker"
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
	var result *models.ExecutionResult
	var err error

	if req.Language == "python" || req.Language == "javascript" || req.Language == "c" || req.Language == "cpp" {
		result, err = dockerrunner.RunCode(
			req.Code,
			lang.Extension,
			lang.DockerInterpreter,
			lang.DockerCompiler,
		)
	} else {
		result, err = sandbox.RunCode(
			req.Code,
			lang.Extension,
			lang.LocalInterpreter,
			lang.LocalCompiler,
		)
	}

	var output string
	if result != nil {
		output = result.Output
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"output": output,
		})
		return
	}

	c.JSON(http.StatusOK, models.ExecuteResponse{
		Output: output,
	})
}
