package handlers

import (
	"net/http"

	"cg-edge-host-api/system"

	"github.com/gin-gonic/gin"
)

//////////////SYSTEM HANDLERS////////////////////

func RestartHostHandler(c *gin.Context) {
	c.JSON(http.StatusOK, system.RestartHost())
}

func ShutDownHostHandler(c *gin.Context) {
	c.JSON(http.StatusOK, system.ShutdownHost())
}
