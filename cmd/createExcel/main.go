package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	// Import interno de packages do próprio sistema
	"product/config"
	"product/pkg/database"
	"product/pkg/entity"
	"product/pkg/service"

	"github.com/xuri/excelize/v2"
)

func main() {
	list_product := ConnectBD().GetAll()

	CreateExcel(list_product)
}

func ConnectBD() *service.Produto_service {
	// Atribui o endereço da estrutura de uma configuração padrão do sistema
	default_conf := &config.Config{}

	// Atribui arquivo Json com configurações externas para file_config, verifica se não está vazia para virar novas configurações
	if file_config := "test_db.json"; file_config != "" {
		file, _ := os.ReadFile(file_config)
		_ = json.Unmarshal(file, &default_conf)
	}

	// Atribui para conf as novas configurações do sistema
	conf := config.NewConfig(default_conf)

	// Pega pool de conexão do Database (config passadas anteriorente pelo Json para Database)
	dbpool := database.NewDB(conf)
	// Se criou uma conexão com as configurações passadas para o Database - Imprima essa mensagem de sucesso
	if dbpool != nil {
		log.Print("Successfully connected")
	}

	// Cria serviços de um produto (CRUD) com a pool de conexão passada por parâmetro
	service := service.NewProdutoService(dbpool)

	return service
}

func CreateExcel(list_product *entity.ProdutoList) {
	f := excelize.NewFile()
	// Set value of a cell.

	styleHeader, err := f.NewStyle(`{
									"border: {
										"type":"right", "color":["#000000"], "style":3}",
									"font":{
										"bold":true},
									"fill":{
										"type":"pattern","pattern":1,"color":["#1976d2"]},
									"alignment":{
										"horizontal":"center","wrap_text":false}}`)
	if err != nil {
		return
	}

	/*styleBold, err := f.NewStyle(`{"font":{"bold":true}}`)
	if err != nil {
		return
	}*/

	/*styleCenter, err := f.NewStyle(`{"alignment":{"horizontal":"center","wrap_text":false}}`)
	if err != nil {
		return
	}*/

	/*stylePrice, err := f.NewStyle(`{"number_format": 353}`)
	if err != nil {
		return
	}*/

	err = f.SetColWidth("Sheet1", "B", "F", 25)
	if err != nil {
		return
	}
	err = f.SetCellStyle("Sheet1", "A1", "F1", styleHeader)
	if err != nil {
		return
	}
	/*err = f.SetRowStyle("Sheet1", 1, 1, styleCenter)
	if err != nil {
		return
	}*/

	/*err = f.SetRowStyle("Sheet1", 2, 1000, styleCenter)
	if err != nil {
		return
	}*/

	/*err = f.SetColStyle("Sheet1", "D", stylePrice)
	if err != nil {
		return
	}*/

	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "Name")
	f.SetCellValue("Sheet1", "C1", "Code")
	f.SetCellValue("Sheet1", "D1", "Price")
	f.SetCellValue("Sheet1", "E1", "Create at")
	f.SetCellValue("Sheet1", "F1", "Update at")

	for i, p := range list_product.List {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), p.ID)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), p.Name)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), p.Code)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+2), p.Price)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+2), p.CreatedAt)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(i+2), p.UpdatedAt)
	}

	// Save spreadsheet by the given path.
	if err := f.SaveAs("Products.xlsx"); err != nil {
		fmt.Println(err)
	}
}
