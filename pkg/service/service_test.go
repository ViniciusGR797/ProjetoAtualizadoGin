package service

import (
	"product/pkg/database"
	"product/pkg/entity"
	"reflect"
	"testing"
)

// Teste unitário do método NewProdutoService
func TestNewProdutoService(t *testing.T) {
	type args struct {
		dabase_pool database.DatabaseInterface
	}
	tests := []struct {
		name string
		args args
		want *Produto_service
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProdutoService(tt.args.dabase_pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProdutoService() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Teste unitário do método GetLog
func TestProduto_service_GetLog(t *testing.T) {
	tests := []struct {
		name string
		ps   *Produto_service
		want *entity.LogList
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.GetLog(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Produto_service.GetLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Teste unitário do método Create
func Test_produto_service_Create(t *testing.T) {
	type args struct {
		produto *entity.Produto
	}
	tests := []struct {
		name string
		ps   *Produto_service
		args args
		want int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.Create(tt.args.produto); got != tt.want {
				t.Errorf("produto_service.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Teste unitário do método Delete
func Test_produto_service_Delete(t *testing.T) {
	type args struct {
		id *int
	}
	tests := []struct {
		name string
		ps   *Produto_service
		args args
		want int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.Delete(tt.args.id); got != tt.want {
				t.Errorf("produto_service.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Teste unitário do método GetAll
func Test_produto_service_GetAll(t *testing.T) {
	tests := []struct {
		name string
		ps   *Produto_service
		want *entity.ProdutoList
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("produto_service.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Teste unitário do método GetProduto
func Test_produto_service_GetProduto(t *testing.T) {
	type args struct {
		ID *int
	}
	tests := []struct {
		name string
		ps   *Produto_service
		args args
		want *entity.Produto
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.GetProduto(tt.args.ID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("produto_service.GetProduto() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Teste unitário do método Update
func Test_produto_service_Update(t *testing.T) {
	type args struct {
		ID      *int
		produto *entity.Produto
	}
	tests := []struct {
		name string
		ps   *Produto_service
		args args
		want int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.Update(tt.args.ID, tt.args.produto); got != tt.want {
				t.Errorf("produto_service.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
