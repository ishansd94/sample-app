package healthz

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ishansd94/sample-app/internal/pkg/response"
)

func Health(c *gin.Context) {
	response.Custom(c, http.StatusOK, gin.H{
		"components": gin.H{
			"app":      "ready",
			"database": "accepted",
		},
	})
	return
}
