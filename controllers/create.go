package controllers

import (
	"product/entity"
	"product/service"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context, service service.ProdutoServiceInterface) {

	var produto *entity.Produto

	err := c.ShouldBind(&produto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON produto" + err.Error(),
		})
		return
	}

	id := service.Create(produto)
	if id == 0 {
		c.JSON(400, gin.H{
			"error": "cannot create JSON: " + err.Error(),
		})
	}

	produto = service.GetProduto(&id)
	c.JSON(200, produto)
}
