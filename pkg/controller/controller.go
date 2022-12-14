package controller

import (
	"product/pkg/entity"
	"product/pkg/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tawesoft/golib/v2/dialog"
)

// Função que chama método GetLog do service (retorna lista de logs) e retorna json com a lista de logs
func GetLog(c *gin.Context, service service.ProdutoServiceInterface) {
	// Chama método GetLog e retorna list de logs
	list := service.GetLog()
	// Verifica se a lista está vazia (tem tamanho zero)
	if len(list.List) == 0 {
		c.JSON(404, gin.H{
			"error": "list not found, 404",
		})
		return
	}
	//retorna sucesso 200 e retorna json da lista de logs
	c.JSON(200, list)
}

// Função que chama método Create do service (retorna id criado) e retorna json com o produto
func Create(c *gin.Context, service service.ProdutoServiceInterface) {
	// Cria variável do tipo produto (inicialmente vazia)
	var produto *entity.Produto

	// Converte json em product
	err := c.ShouldBind(&produto)
	// Verifica se tem erro
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON product" + err.Error(),
		})
		return
	}

	// Verifica se o price é positivo (não aceita valor negativo e igual a zero)
	if produto.Price <= 0.0 {
		c.JSON(400, gin.H{
			"error": "Invalid value in price",
		})

		// Dispara pop up de erro em valor em price
		dialog.Alert("Invalid value in price")
		return
	}

	// Chama método Create passando produto como parâmetro que retorna id novo
	id := service.Create(produto)
	// Verifica se o id é zero (caso for deu erro ao criar produto no banco)
	if id == 0 {
		c.JSON(400, gin.H{
			"error": "cannot create JSON: " + err.Error(),
		})
	}

	// Pega produto no banco
	produto = service.GetProduto(&id)
	// Retorno json com o produto
	c.JSON(200, produto)
}

// Função que chama método Delete do service e retorna json com mensagem de deletado
func Delete(c *gin.Context, service service.ProdutoServiceInterface) {
	// Pega id passada como parâmetro na URL da rota
	id := c.Param("id")

	// Converter ":id" string para int id (newid)
	newId, err := strconv.Atoi(strings.Replace(id, ":", "", 1))
	// Verifica se teve erro na conversão
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400",
		})
		return
	}

	// Chama método Delete passando id como parâmetro
	aff := service.Delete(&newId)
	// Verifica se teve erro ao deletar produto
	if aff == 0 {
		c.JSON(400, gin.H{
			"error": "cannot delete produto, 400",
		})
		return
	}

	// Retorno json com mensagem de deletado
	c.JSON(200, gin.H{
		"mensage": "Produto deleted",
	})
}

// Função que chama método GetAll do service e retorna json com lista de produtos
func GetAll(c *gin.Context, service service.ProdutoServiceInterface) {
	// Chama método GetAll e retorna list de products
	list := service.GetAll()
	// Verifica se a lista está vazia (tem tamanho zero)
	if len(list.List) == 0 {
		c.JSON(404, gin.H{
			"error": "lista not found, 404",
		})
		return
	}
	//retorna sucesso 200 e retorna json da lista de products
	c.JSON(200, list)
}

// Função que chama método GetProduto do service e retorna json com produto
func GetProduto(c *gin.Context, service service.ProdutoServiceInterface) {
	// Pega id passada como parâmetro na URL da rota
	id := c.Param("id")

	// Converter ":id" string para int id (newid)
	newId, err := strconv.Atoi(strings.Replace(id, ":", "", 1))
	// Verifica se teve erro na conversão
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400",
		})
		return
	}
	// Chama método GetProduto passando id como parâmetro
	produto := service.GetProduto(&newId)
	if produto.ID == 0 {
		c.JSON(404, gin.H{
			"error": "product not found, 404",
		})
		return
	}

	// Retorno json com produto
	c.JSON(200, produto)
}

// Função que chama método Update do service e retorna json com produto alterado
func Update(c *gin.Context, service service.ProdutoServiceInterface) {
	// Pega id passada como parâmetro na URL da rota
	id := c.Param("id")
	// Cria variável do tipo produto (inicialmente vazia)
	var produto *entity.Produto

	// Converter ":id" string para int id (newid)
	newId, err := strconv.Atoi(strings.Replace(id, ":", "", 1))
	// Verifica se teve erro na conversão
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400" + err.Error(),
		})
		return
	}

	// Converte json em product
	err = c.ShouldBind(&produto)
	// Verifica se tem erro
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON produto, 400" + err.Error(),
		})
		return
	}

	// Verifica se o price é positivo (não aceita valor negativo e igual a zero)
	if produto.Price <= 0.0 {
		c.JSON(400, gin.H{
			"error": "invalid value in price",
		})

		// Dispara pop up de erro em valor em price
		dialog.Alert("Invalid value in price")
		return
	}

	// Chama método Update passando produto e id editado como parâmetro
	idResult := service.Update(&newId, produto)
	// Verifica se o id é zero (caso for deu erro ao editar produto no banco)
	if idResult == 0 {
		c.JSON(400, gin.H{
			"error": "cannot update JSON, 400" + err.Error(),
		})
		return
	}

	// Pega produto no banco
	produto = service.GetProduto(&newId)
	// Retorno json com o produto
	c.JSON(200, produto)
}
