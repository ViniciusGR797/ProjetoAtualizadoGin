package controllers

import (
	"product/entity"
	"product/service"

	"github.com/gin-gonic/gin"
	"github.com/tawesoft/golib/v2/dialog"
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

	if produto.Price <= 0.0 {
		c.JSON(400, gin.H{
			"error": "Invalid value in price",
		})

		dialog.Alert("Invalid value in price")
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
