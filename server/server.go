package server

import (
	"log"

	// Import interno de packages do próprio sistema
	"product/config"
	"product/service"

	// Import externo do github
	"github.com/gin-gonic/gin"
)

// Estrutura de dados para armazenar o servidor HTTP
type Server struct {
	// Porta do servidor
	SRV_PORT string

	// Ponteiro de servidor do framework gin
	SERVER *gin.Engine
}

// Cria novo servidor HTTP, de acordo com as config passadas por parâmetro
func NewServer(conf *config.Config) Server {
	return Server{
		SRV_PORT: conf.SRV_PORT,
		SERVER:   gin.Default(),
	}
}

// Rodar servidor HTTP, tendo as rotas do framework gin, servidor HTTP, serviço CRUD de produto
func Run(router *gin.Engine, server Server, service service.ProdutoServiceInterface) {
	// Imprime que servidor HTTP está rodando na porta tal
	log.Print("Server is running at port: ", server.SRV_PORT)

	/*err := router.Run(":" + server.SRV_PORT)
	if err != nil {
		log.Fatal(err)
	}*/

	log.Fatal(router.Run("0.0.0.0:" + server.SRV_PORT))
}
