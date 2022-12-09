package controllers

import (
	"api-produto/service"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context, service service.ProdutoServiceInterface) {

	lista := service.GetAll()
	if len(lista.List) == 0 {
		c.JSON(404, gin.H{
			"error": "lista not found, 404",
		})
		return
	}
	c.JSON(200, lista)
}
