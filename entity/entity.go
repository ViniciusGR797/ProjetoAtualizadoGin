package entity

import (
	"encoding/json"
	"log"
)

type ProdutoInterface interface {
	toString() string
}

type Produto struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Price       float64 `json:"price"`
	CriadoEm    string  `json:"criado_em,omitempty"`
	AtulizadoEm string  `json:"atulizado_em,omitempty"`
}

func (p Produto) toString() string {
	data, err := json.Marshal(p)
	if err != nil {
		log.Printf("Erro ao converter Produto para String: \n--> %v", err)
	}

	return string(data)
}

func NovoProduto(nome, code string, price float64) *Produto {
	return &Produto{
		Name:  nome,
		Code:  code,
		Price: price,
	}
}

type ListaDeProduto struct {
	List []*Produto `json:"lista"`
}

func (prodLista ListaDeProduto) toString() string {
	data, err := json.Marshal(prodLista)
	if err != nil {
		log.Printf("Erro ao converter Produto para String: \n--> %v", err)
	}

	return string(data)
}

type User struct {
	Username string `json:"username"`
	Senha    string `json:"senha"`
}

func NovoAdmin() *User {
	return &User{
		Username: "admin",
		Senha:    "supersenha",
	}
}

type Token struct {
	Token string `json:"token"`
}

const USER_TOKEN = "fake-WzD5fqrlaAXLv26bpI0hxvAhDp7T1Bac"
