package controller

import (
	"product/pkg/service"
	"testing"

	"github.com/gin-gonic/gin"
)

// Teste unitário do método GetLog
func TestGetLog(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ProdutoServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetLog(tt.args.c, tt.args.service)
		})
	}
}

// Teste unitário do método Create
func TestCreate(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ProdutoServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Create(tt.args.c, tt.args.service)
		})
	}
}

// Teste unitário do método Delete
func TestDelete(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ProdutoServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Delete(tt.args.c, tt.args.service)
		})
	}
}

// Teste unitário do método GetAll
func TestGetAll(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ProdutoServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetAll(tt.args.c, tt.args.service)
		})
	}
}

// Teste unitário do método GetProduto
func TestGetProduto(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ProdutoServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetProduto(tt.args.c, tt.args.service)
		})
	}
}

// Teste unitário do método Update
func TestUpdate(t *testing.T) {
	type args struct {
		c       *gin.Context
		service service.ProdutoServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Update(tt.args.c, tt.args.service)
		})
	}
}
