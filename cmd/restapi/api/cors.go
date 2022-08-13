package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AllowAllOrigins() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	return cors.New(corsConfig)
}
