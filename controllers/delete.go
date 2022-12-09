package controllers

import (
	"api-produto/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context, service service.ProdutoServiceInterface) {

	id := c.Param("id")

	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400",
		})
		return
	}

	aff := service.Delete(&newId)
	if aff == 0 {
		c.JSON(400, gin.H{
			"error": "cannot delete produto, 400",
		})
		return
	}

	c.JSON(200, gin.H{
		"mensage": "Produto deleted",
	})
}
