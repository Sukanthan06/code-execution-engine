package handlers

import (
	"net/http"

	dockerrunner "github.com/Sukanthan06/code-execution-engine/executor/docker"
	"github.com/Sukanthan06/code-execution-engine/executor/sandbox"
	"github.com/Sukanthan06/code-execution-engine/languages"
	"github.com/Sukanthan06/code-execution-engine/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ExecuteCode(c *gin.Context) {
	executionID := uuid.New().String()

	var req models.ExecuteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ExecuteResponse{
			ExecutionID: executionID,
			Status:      models.StatusInternalError,
			Error:       err.Error(),
		})
		return
	}

	lang, exists := languages.GetLanguage(req.Language)

	if !exists {
		c.JSON(http.StatusBadRequest, models.ExecuteResponse{
			ExecutionID: executionID,
			Status:      models.StatusUnsupportedLanguage,
			Error:       "unsupported language",
		})
		return
	}
	var result *models.ExecutionResult
	var err error

	if lang.DockerInterpreter != "" || lang.DockerCompiler != "" {
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
		errMsg := err.Error()
		if result != nil && result.Error != "" {
			errMsg = result.Error
		}
		resp := models.ExecuteResponse{
			ExecutionID: executionID,
			Output:      output,
			Error:       errMsg,
		}
		if result != nil {
			resp.Status = result.Status
			resp.ExitCode = result.ExitCode
			resp.RuntimeMS = result.RuntimeMS
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, models.ExecuteResponse{
		ExecutionID: executionID,
		Output:      output,
		Status:      result.Status,
		ExitCode:    result.ExitCode,
		RuntimeMS:   result.RuntimeMS,
	})
}
