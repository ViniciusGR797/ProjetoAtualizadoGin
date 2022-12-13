package service

import (
	"product/pkg/database"
	"product/pkg/entity"
	"reflect"
	"testing"
)

func TestNewProdutoService(t *testing.T) {
	type args struct {
		dabase_pool database.DatabaseInterface
	}
	tests := []struct {
		name string
		args args
		want *produto_service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProdutoService(tt.args.dabase_pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProdutoService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_produto_service_Create(t *testing.T) {
	type args struct {
		produto *entity.Produto
	}
	tests := []struct {
		name string
		ps   *produto_service
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.Create(tt.args.produto); got != tt.want {
				t.Errorf("produto_service.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_produto_service_Delete(t *testing.T) {
	type args struct {
		id *int
	}
	tests := []struct {
		name string
		ps   *produto_service
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.Delete(tt.args.id); got != tt.want {
				t.Errorf("produto_service.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_produto_service_GetAll(t *testing.T) {
	tests := []struct {
		name string
		ps   *produto_service
		want *entity.ProdutoList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("produto_service.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_produto_service_GetProduto(t *testing.T) {
	type args struct {
		ID *int
	}
	tests := []struct {
		name string
		ps   *produto_service
		args args
		want *entity.Produto
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.GetProduto(tt.args.ID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("produto_service.GetProduto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_produto_service_Update(t *testing.T) {
	type args struct {
		ID      *int
		produto *entity.Produto
	}
	tests := []struct {
		name string
		ps   *produto_service
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.Update(tt.args.ID, tt.args.produto); got != tt.want {
				t.Errorf("produto_service.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
