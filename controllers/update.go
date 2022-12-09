package controllers

import (
	"api-produto/entity"
	"api-produto/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context, service service.ProdutoServiceInterface) {

	id := c.Param("id")

	var produto *entity.Produto

	newId, err := strconv.ParseInt(id, 10, 64)
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
