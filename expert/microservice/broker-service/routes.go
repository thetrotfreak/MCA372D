package main

import "github.com/gin-gonic/gin"

// initializing a router
func (app *Config) Routes() *gin.Engine {
	router := gin.Default()
	group := router.Group("api/v1/broker-service")
	group.GET("/ping", app.Ping)

	return router
}
