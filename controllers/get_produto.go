package controllers

import (
	"product/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProduto(c *gin.Context, service service.ProdutoServiceInterface) {

	id := c.Param("id")

	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400",
		})
		return
	}

	produto := service.GetProduto(&newId)
	if produto.ID == 0 {
		c.JSON(404, gin.H{
			"error": "produto not found, 404",
		})
		return
	}

	c.JSON(200, produto)

}
