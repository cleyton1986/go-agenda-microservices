package main

import (
	"github.com/cleyton1986/go-agenda-microservices/cmd/internal/contacts/controllers"
	"github.com/gin-gonic/gin"
)

func ContactRoutes(router *gin.Engine){

	contactsRoutes := router.Group("/contacts")
	{
		contactsRoutes.POST("/", controllers.CreateContact)
	}


}
