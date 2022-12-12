package controllers

import (
	"product/entity"
	"product/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tawesoft/golib/v2/dialog"
)

func Update(c *gin.Context, service service.ProdutoServiceInterface) {

	id := c.Param("id")

	var produto *entity.Produto

	newId, err := strconv.Atoi(strings.Replace(id, ":", "", 1))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400" + err.Error(),
		})
		return
	}

	err = c.ShouldBind(&produto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON produto, 400" + err.Error(),
		})
		return
	}

	if produto.Price <= 0.0 {
		c.JSON(400, gin.H{
			"error": "invalid value in price",
		})

		dialog.Alert("Invalid value in price")
		return
	}

	idResult := service.Update(&newId, produto)
	if idResult == 0 {
		c.JSON(400, gin.H{
			"error": "cannot update JSON, 400" + err.Error(),
		})
		return
	}

	produto = service.GetProduto(&newId)
	c.JSON(200, produto)
}
