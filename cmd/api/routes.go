package main

import (
	"jwt-gin/controllers"
	"jwt-gin/middlewares"
	"jwt-gin/models"

	"github.com/gin-gonic/gin"
)

/* ====================================================
#   Copyright (C) 2022  All rights reserved.
#
#   Author        : wander
#   Email         : wander@ffactory.org
#   File Name     : routes.go
#   Last Modified : 2022-08-30 12:00
#   Describe      :
#
# ====================================================*/

func routes() *gin.Engine {
	models.ConnectDataBase()
	r := gin.Default()
	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)
	r.Run(":8080")
	return r
}
