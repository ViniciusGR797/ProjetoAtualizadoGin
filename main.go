package main

import (
	"encoding/json"
	"log"
	"os"

	"product/config"
	"product/database"
	"product/routes"
	"product/server"
	"product/service"
	"product/webui"
)

func main() {

	default_conf := &config.Config{}

	if file_config := "test_db.json"; file_config != "" {
		file, _ := os.ReadFile(file_config)
		_ = json.Unmarshal(file, &default_conf)
	}

	conf := config.NewConfig(default_conf)

	dbpool := database.NewDB(conf)
	if dbpool != nil {
		log.Print("Successfully connected: ", dbpool.GetDB())
	}

	service := service.NewProdutoService(dbpool)
	serv := server.NewServer(conf)

	// Inicia as rotas
	router := routes.ConfigRoutes(serv.SERVER, service)

	if conf.WEB_UI {
		webui.RegisterUIHandlers(router)
	}

	// Registra as rotas
	server.Run(router, serv, service)

}
