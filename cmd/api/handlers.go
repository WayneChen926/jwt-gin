package main

import (
	"net/http"

	"jwt-gin/config"

	"github.com/gin-gonic/gin"
)

/* ====================================================
#   Copyright (C) 2022  All rights reserved.
#
#   Author        : wander
#   Email         : wander@ffactory.org
#   File Name     : handlers.go
#   Last Modified : 2022-08-30 12:06
#   Describe      :
#
# ====================================================*/

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func ping(c *gin.Context) {
	getString := app.SetString
	c.JSON(http.StatusOK, gin.H{"status": getString})
}
