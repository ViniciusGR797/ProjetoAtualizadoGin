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

	styleHeader, err := f.NewStyle(`{"Border: {"Type":"Continuous", "Color":["#000000"], "Style":2}", "fill":{"type":"pattern","pattern":1,"color":["#1976d2"]},"font":{"bold":true},"alignment":{"horizontal":"center","ident":1,"justify_last_line":true,"reading_order":0,"relative_indent":1,"shrink_to_fit":true,"vertical":"","wrap_text":true}}`)
	if err != nil {
		return
	}

	styleBody, err := f.NewStyle(`{"Border: {"Type":"Continuous", "Color":["#000000"], "Style":2}", "fill":{"type":"pattern","pattern":1,"color":["#eeeeee"]},"alignment":{"horizontal":"center","ident":1,"justify_last_line":true,"reading_order":0,"relative_indent":1,"shrink_to_fit":true,"vertical":"","wrap_text":false}}`)
	if err != nil {
		return
	}

	err = f.SetColWidth("Sheet1", "B", "F", 25)
	if err != nil {
		return
	}

	/*err = f.SetPageMargins("Sheet1", "Bottom")
	if err != nil {
		return
	}*/

	err = f.SetRowStyle("Sheet1", 1, 1, styleHeader)
	if err != nil {
		return
	}

	err = f.SetRowStyle("Sheet1", 2, 1000, styleBody)
	if err != nil {
		return
	}

	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "Name")
	f.SetCellValue("Sheet1", "C1", "Code")
	f.SetCellValue("Sheet1", "D1", "Price")
	f.SetCellValue("Sheet1", "E1", "Create at")
	f.SetCellValue("Sheet1", "F1", "Update at")

	for i, p := range list_product.List {
		//s, _ :=

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
