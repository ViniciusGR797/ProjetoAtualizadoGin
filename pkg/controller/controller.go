package controller

import (
	"product/pkg/entity"
	"product/pkg/service"
	"strconv"
	"strings"

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

func Delete(c *gin.Context, service service.ProdutoServiceInterface) {

	id := c.Param("id")

	newId, err := strconv.Atoi(strings.Replace(id, ":", "", 1))

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

func GetProduto(c *gin.Context, service service.ProdutoServiceInterface) {

	id := c.Param("id")

	newId, err := strconv.Atoi(strings.Replace(id, ":", "", 1))

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400",
		})
		return
	}

	produto := service.GetProduto(&newId)
	if produto.ID == 0 {
		c.JSON(404, gin.H{
			"error": "product not found, 404",
		})
		return
	}

	c.JSON(200, produto)
}

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
