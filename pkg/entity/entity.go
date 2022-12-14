package entity

import (
	"encoding/json"
	"log"
)

type ProdutoInterface interface {
	String() string
}

type LogInterface interface {
	String() string
}

// Estrutura de dados de Produto
type Produto struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Code      string  `json:"code"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at,omitempty"`
	UpdatedAt string  `json:"updated_at,omitempty"`
	//gorm.Model
}

// Método de produto - retorna string com json do produto ou erro
func (p *Produto) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert Produto to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// Estrutura de dados para lista de Produtos
type ProdutoList struct {
	List []*Produto `json:"list"`
}

// Método de ProdutoList - retorna string com json da lista de produtos ou erro
func (pl *ProdutoList) String() string {
	data, err := json.Marshal(pl)

	if err != nil {
		log.Println("error to convert ProdutoList to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// Construtor de Produto - recebe dados no parâmetro e transforma em um produto
func NewProduto(nome, code string, price float64) *Produto {
	return &Produto{
		Name:  nome,
		Code:  code,
		Price: price,
	}
}

// Estrutura de dados para user - usuário que usará o sistema
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Construtor de User - terá o admin como usuário padrão do sistema
func NewAdmin() *User {
	return &User{
		Username: "grupob",
		Password: "grupob",
	}
}

// Estrutura de dados de Log
type Log struct {
	ID          int    `json:"id"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Data        string `json:"data,omitempty"`
	//gorm.Model
}

// Estrutura de dados para lista de Log
type LogList struct {
	List []*Log `json:"list"`
}

// Método de LogList - retorna string com json da lista de log ou erro
func (ll *LogList) String() string {
	data, err := json.Marshal(ll)

	if err != nil {
		log.Println("error to convert LogList to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// Estrutura de dados para token - key de acesso ao sistema para estar autenticado
type Token struct {
	Token string `json:"token"`
}

// Token do user - trata de um token fake, pois ele é uma constante que não altera
const USER_TOKEN = "fake-WzD5fqrlaAXLv26bpI0hxvAhDp7T1Bac"
