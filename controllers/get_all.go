package controllers

import (
	"product/service"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context, service service.ProdutoServiceInterface) {

	list := service.GetAll()
	if len(list.List) == 0 {
		c.JSON(404, gin.H{
			"error": "lista not found, 404",
		})
		return
	}
	c.JSON(200, list)
}
