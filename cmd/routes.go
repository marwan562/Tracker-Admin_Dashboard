package main

import "github.com/gin-gonic/gin"

func setupRoutes(r *gin.Engine, handler *Handler) {
	r.GET("/", handler.ServeNewOrder)
	r.POST("/new-order", handler.HandleNewOrderPost)
	r.GET("/customer/:id", handler.ServeCustomer)

	r.Static("/static", "./templates/static")
}
