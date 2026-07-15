package handlers

import (
	"net/http"

	dockerrunner "github.com/Sukanthan06/code-execution-engine/executor/docker"
	"github.com/gin-gonic/gin"
)

func DockerTest(c *gin.Context) {
	output, err := dockerrunner.TestDocker()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"output": output,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"output": output,
	})
}
