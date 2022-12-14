package routes

import (
	"product/pkg/controller"
	"product/pkg/service"

	"github.com/gin-gonic/gin"
)

// Função que configura todas as rotas da api
func ConfigRoutes(router *gin.Engine, service service.ProdutoServiceInterface) *gin.Engine {
	main := router.Group("api")
	{
		produtos := main.Group("/v1")
		{
			// Rota que retorna lista de log (GET que dispara método GetLog controller)
			produtos.GET("/logs", func(c *gin.Context) {
				controller.GetLog(c, service)
			})
			// Rota que retorna product (GET que dispara método GetProduto controller)
			produtos.GET("/product/:id", func(c *gin.Context) {
				controller.GetProduto(c, service)
			})
			// Rota que retorna lista de product (GET que dispara método GetAll controller)
			produtos.GET("/products", func(c *gin.Context) {
				controller.GetAll(c, service)
			})
			// Rota que retorna product criado (POST que dispara método Create controller)
			produtos.POST("/product", func(c *gin.Context) {
				controller.Create(c, service)
			})
			// Rota que retorna product editado (PUT que dispara método Update controller)
			produtos.PUT("/product/:id", func(c *gin.Context) {
				controller.Update(c, service)
			})
			// Rota que retorna mensagem de deletado (DELETE que dispara método Delete controller)
			produtos.DELETE("/product/:id", func(c *gin.Context) {
				controller.Delete(c, service)
			})
			// Rota para login de usuário
			produtos.POST("/user/login", func(c *gin.Context) {
				c.Next()
			})
		}
	}

	// retorna rota
	return router
}
