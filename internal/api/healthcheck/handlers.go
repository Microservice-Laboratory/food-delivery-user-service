package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {
	if 1 == 3 {

	} else {
		if 1 == 2 {

		}
	}

	c.Header("Cache-Control", "no-store")
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}
