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
	list_log := ConnectBD().GetLog()

	CreateExcel(list_log)
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

func CreateExcel(list_log *entity.LogList) {
	f := excelize.NewFile()
	// Set value of a cell.

	styleHeader, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 2},
			{Type: "top", Color: "#000000", Style: 2},
			{Type: "bottom", Color: "#000000", Style: 2},
			{Type: "right", Color: "#000000", Style: 2},
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#6959CD"},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			WrapText:   true,
		},
		Font: &excelize.Font{
			Bold:  true,
			Color: "#FFFFFF",
		},
	})
	if err != nil {
		return
	}

	styleBodyGET, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 2},
			{Type: "top", Color: "#000000", Style: 2},
			{Type: "bottom", Color: "#000000", Style: 2},
			{Type: "right", Color: "#000000", Style: 2},
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#808080"},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			WrapText:   true,
		},
		Font: &excelize.Font{
			Color: "#00FFFF",
		},
	})
	if err != nil {
		return
	}

	styleBodyPOST, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 2},
			{Type: "top", Color: "#000000", Style: 2},
			{Type: "bottom", Color: "#000000", Style: 2},
			{Type: "right", Color: "#000000", Style: 2},
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#808080"},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			WrapText:   true,
		},
		Font: &excelize.Font{
			Color: "#00FF7F",
		},
	})
	if err != nil {
		return
	}

	styleBodyPUT, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 2},
			{Type: "top", Color: "#000000", Style: 2},
			{Type: "bottom", Color: "#000000", Style: 2},
			{Type: "right", Color: "#000000", Style: 2},
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#808080"},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			WrapText:   true,
		},
		Font: &excelize.Font{
			Color: "#EE82EE",
		},
	})
	if err != nil {
		return
	}

	styleBodyDELETE, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 2},
			{Type: "top", Color: "#000000", Style: 2},
			{Type: "bottom", Color: "#000000", Style: 2},
			{Type: "right", Color: "#000000", Style: 2},
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#808080"},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			WrapText:   true,
		},
		Font: &excelize.Font{
			Color: "#B22222",
		},
	})
	if err != nil {
		return
	}

	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "Method")
	f.SetCellValue("Sheet1", "C1", "Description")
	f.SetCellValue("Sheet1", "D1", "Data")

	for i, l := range list_log.List {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), l.ID)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), l.Method)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), l.Description)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+2), l.Data)

		switch l.Method {
		case "GET":
			err = f.SetCellStyle("Sheet1", "A"+strconv.Itoa(i+2), "D"+strconv.Itoa(i+2), styleBodyGET)
			if err != nil {
				return
			}
		case "POST":
			err = f.SetCellStyle("Sheet1", "A"+strconv.Itoa(i+2), "D"+strconv.Itoa(i+2), styleBodyPOST)
			if err != nil {
				return
			}
		case "PUT":
			err = f.SetCellStyle("Sheet1", "A"+strconv.Itoa(i+2), "D"+strconv.Itoa(i+2), styleBodyPUT)
			if err != nil {
				return
			}
		case "DELETE":
			err = f.SetCellStyle("Sheet1", "A"+strconv.Itoa(i+2), "D"+strconv.Itoa(i+2), styleBodyDELETE)
			if err != nil {
				return
			}
		}

	}

	err = f.SetColWidth("Sheet1", "C", "C", 50)
	if err != nil {
		return
	}
	err = f.SetColWidth("Sheet1", "D", "D", 25)
	if err != nil {
		return
	}
	err = f.SetCellStyle("Sheet1", "A1", "D1", styleHeader)
	if err != nil {
		return
	}

	// Save spreadsheet by the given path.
	if err := f.SaveAs("Logs.xlsx"); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Excel of logs successfully generated")
}
