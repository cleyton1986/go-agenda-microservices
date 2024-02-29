package controllers

import (
	"net/http"
	"time"

	use_case "github.com/cleyton1986/go-agenda-microservices/cmd/internal/contacts/use-case"
	"github.com/gin-gonic/gin"
)

type createContactInput struct{
	Name      string    `json:"name" binding:"required"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateContact(ctx *gin.Context) {

	var body createContactInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"sucess": false,
			"error": err.Error(),
		})


		return
	}

	useCase := use_case.NewCreateContactUseCase()

	err := useCase.Execute(body.Name, body.Surname, body.Email)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"sucess": false,
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"sucess":true})
}