package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *Config) Ping(ctx *gin.Context) {
	log.Println("PING to server")
	ctx.JSON(http.StatusOK, "Message: pong")
}
