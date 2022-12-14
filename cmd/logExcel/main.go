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

// Função principal (primeira executada) - chama métodos que vão config para fazer conexão BD, service e criar planilha
func main() {
	// Chama método ConnectBD() para config para fazer conexão BD e service, já o GetLog() retorna Lista dos logs para uma variável
	list_log := ConnectBD().GetLog()

	// Chama método CreateExcel() que vai criar planilha Excel com a list passada por parâmetro
	CreateExcel(list_log)
}

// Função que vai configurar conexão BD e service, retorna service criado
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

	// Retorna service válido
	return service
}

// Função que cria a planilha Excel com a list passada por parâmetro
func CreateExcel(list_log *entity.LogList) {
	// Cria arquivo de planilha Excel vazia e atribui a variável f
	f := excelize.NewFile()

	// Cria novo estilo para o cabeçalho da planilha (formatando planilha)
	styleHeader, err := f.NewStyle(&excelize.Style{
		// Borda completa na planilha (left, top, bottom e right)
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 2},
			{Type: "top", Color: "#000000", Style: 2},
			{Type: "bottom", Color: "#000000", Style: 2},
			{Type: "right", Color: "#000000", Style: 2},
		},
		// Preenchimento da célula (deixa colorida)
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#6959CD"},
		},
		// Alinhamneto do texto na célula (centralizado e não dar quebra de linha)
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			WrapText:   true,
		},
		// Fonte em negrito e colorida
		Font: &excelize.Font{
			Bold:  true,
			Color: "#FFFFFF",
		},
	})
	// Verifica se teve algum erro ao criar esse novo estilo
	if err != nil {
		return
	}

	// Cria novo estilo para o corpo para method GET da planilha (formatando planilha)
	styleBodyGET, err := f.NewStyle(&excelize.Style{
		// Borda completa na planilha (left, top, bottom e right)
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 2},
			{Type: "top", Color: "#000000", Style: 2},
			{Type: "bottom", Color: "#000000", Style: 2},
			{Type: "right", Color: "#000000", Style: 2},
		},
		// Preenchimento da célula (deixa colorida)
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#FFFFFF"},
		},
		// Alinhamneto do texto na célula (centralizado e não dar quebra de linha)
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			WrapText:   true,
		},
		// Fonte em negrito e colorida
		Font: &excelize.Font{
			Bold:  true,
			Color: "#000000",
		},
	})
	// Verifica se teve algum erro ao criar esse novo estilo
	if err != nil {
		return
	}

	// Cria novo estilo para o corpo para method POST da planilha (formatando planilha)
	styleBodyPOST, err := f.NewStyle(&excelize.Style{
		// Borda completa na planilha (left, top, bottom e right)
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 2},
			{Type: "top", Color: "#000000", Style: 2},
			{Type: "bottom", Color: "#000000", Style: 2},
			{Type: "right", Color: "#000000", Style: 2},
		},
		// Preenchimento da célula (deixa colorida)
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#FFFFFF"},
		},
		// Alinhamneto do texto na célula (centralizado e não dar quebra de linha)
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			WrapText:   true,
		},
		// Fonte em negrito e colorida
		Font: &excelize.Font{
			Bold:  true,
			Color: "#006400",
		},
	})
	// Verifica se teve algum erro ao criar esse novo estilo
	if err != nil {
		return
	}

	// Cria novo estilo para o corpo para method PUT da planilha (formatando planilha)
	styleBodyPUT, err := f.NewStyle(&excelize.Style{
		// Borda completa na planilha (left, top, bottom e right)
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 2},
			{Type: "top", Color: "#000000", Style: 2},
			{Type: "bottom", Color: "#000000", Style: 2},
			{Type: "right", Color: "#000000", Style: 2},
		},
		// Preenchimento da célula (deixa colorida)
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#FFFFFF"},
		},
		// Alinhamneto do texto na célula (centralizado e não dar quebra de linha)
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			WrapText:   true,
		},
		// Fonte em negrito e colorida
		Font: &excelize.Font{
			Bold:  true,
			Color: "#003399",
		},
	})
	// Verifica se teve algum erro ao criar esse novo estilo
	if err != nil {
		return
	}

	// Cria novo estilo para o corpo para method PUT da planilha (formatando planilha)
	styleBodyDELETE, err := f.NewStyle(&excelize.Style{
		// Borda completa na planilha (left, top, bottom e right)
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 2},
			{Type: "top", Color: "#000000", Style: 2},
			{Type: "bottom", Color: "#000000", Style: 2},
			{Type: "right", Color: "#000000", Style: 2},
		},
		// Preenchimento da célula (deixa colorida)
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#FFFFFF"},
		},
		// Alinhamneto do texto na célula (centralizado e não dar quebra de linha)
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			WrapText:   true,
		},
		// Fonte em negrito e colorida
		Font: &excelize.Font{
			Bold:  true,
			Color: "#B22222",
		},
	})
	// Verifica se teve algum erro ao criar esse novo estilo
	if err != nil {
		return
	}

	// Colocando os dados no cabeçalho da planilha
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "Method")
	f.SetCellValue("Sheet1", "C1", "Description")
	f.SetCellValue("Sheet1", "D1", "Data")

	// Percorrendo e colocando os dados do corpo da planilha (através da lista de logs)
	for i, l := range list_log.List {
		// Colocando os dados do corpo da planilha (lilha por lilha)
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), l.ID)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), l.Method)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), l.Description)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+2), l.Data)

		// Switch case para verificar qual method http é para colocar o estilo de formatação correto
		switch l.Method {
		case "GET":
			// Colocando formatação de estilo no corpo com method GET
			err = f.SetCellStyle("Sheet1", "A"+strconv.Itoa(i+2), "D"+strconv.Itoa(i+2), styleBodyGET)
			// Verifica se teve algum erro ao colocar estilo
			if err != nil {
				return
			}
		case "POST":
			// Colocando formatação de estilo no corpo com method POST
			err = f.SetCellStyle("Sheet1", "A"+strconv.Itoa(i+2), "D"+strconv.Itoa(i+2), styleBodyPOST)
			// Verifica se teve algum erro ao colocar estilo
			if err != nil {
				return
			}
		case "PUT":
			// Colocando formatação de estilo no corpo com method PUT
			err = f.SetCellStyle("Sheet1", "A"+strconv.Itoa(i+2), "D"+strconv.Itoa(i+2), styleBodyPUT)
			// Verifica se teve algum erro ao colocar estilo
			if err != nil {
				return
			}
		case "DELETE":
			// Colocando formatação de estilo no corpo com method DELETE
			err = f.SetCellStyle("Sheet1", "A"+strconv.Itoa(i+2), "D"+strconv.Itoa(i+2), styleBodyDELETE)
			// Verifica se teve algum erro ao colocar estilo
			if err != nil {
				return
			}
		}

	}

	// Colocando tamanho na coluna C de 50
	err = f.SetColWidth("Sheet1", "C", "C", 50)
	// Verifica se teve algum erro ao colocar tamanho
	if err != nil {
		return
	}
	// Colocando tamanho na coluna D de 25
	err = f.SetColWidth("Sheet1", "D", "D", 25)
	// Verifica se teve algum erro ao colocar tamanho
	if err != nil {
		return
	}
	// Colocando formatação de estilo no cabeçalho
	err = f.SetCellStyle("Sheet1", "A1", "D1", styleHeader)
	// Verifica se teve algum erro ao colocar estilo
	if err != nil {
		return
	}

	// Salva a planilha Excel e coloca nome Logs.xlsx
	if err := f.SaveAs("Logs.xlsx"); err != nil {
		fmt.Println(err)
	}

	// Printa que foi gerado o Excel com sucesso
	fmt.Println("Excel of logs successfully generated")
}
