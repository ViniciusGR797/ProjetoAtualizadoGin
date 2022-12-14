package routes

import (
	"product/pkg/controller"
	"product/pkg/service"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine, service service.ProdutoServiceInterface) *gin.Engine {
	main := router.Group("api")
	{
		produtos := main.Group("/v1")
		{
			produtos.GET("/logs", func(c *gin.Context) {
				controller.GetLog(c, service)
			})
			produtos.GET("/product/:id", func(c *gin.Context) {
				controller.GetProduto(c, service)
			})
			produtos.GET("/products", func(c *gin.Context) {
				controller.GetAll(c, service)
			})
			produtos.POST("/product", func(c *gin.Context) {
				controller.Create(c, service)
			})
			produtos.PUT("/product/:id", func(c *gin.Context) {
				controller.Update(c, service)
			})
			produtos.DELETE("/product/:id", func(c *gin.Context) {
				controller.Delete(c, service)
			})
			produtos.POST("/user/login", func(c *gin.Context) {
				c.Next()
			})
		}
	}

	return router
}
