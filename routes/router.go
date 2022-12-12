package routes

import (
	"product/controllers"
	"product/service"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine, service service.ProdutoServiceInterface) *gin.Engine {
	main := router.Group("api")
	{
		produtos := main.Group("/v1")
		{
			produtos.GET("/product/:id", func(c *gin.Context) {
				controllers.GetProduto(c, service)
			})
			produtos.GET("/products", func(c *gin.Context) {
				controllers.GetAll(c, service)
			})
			produtos.POST("/product", func(c *gin.Context) {
				controllers.Create(c, service)
			})
			produtos.PUT("/product/:id", func(c *gin.Context) {
				controllers.Update(c, service)
			})
			produtos.DELETE("/product/:id", func(c *gin.Context) {
				controllers.Delete(c, service)
			})
			produtos.POST("/user/login", func(c *gin.Context) {
				c.Next()
			})
		}
	}

	return router
}
