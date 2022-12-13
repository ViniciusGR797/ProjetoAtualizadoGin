package entity

import (
	"reflect"
	"testing"
)

func TestProduto_String(t *testing.T) {
	tests := []struct {
		name string
		p    *Produto
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Produto.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProdutoList_String(t *testing.T) {
	tests := []struct {
		name string
		pl   *ProdutoList
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pl.String(); got != tt.want {
				t.Errorf("ProdutoList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewProduto(t *testing.T) {
	type args struct {
		nome  string
		code  string
		price float64
	}
	tests := []struct {
		name string
		args args
		want *Produto
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProduto(tt.args.nome, tt.args.code, tt.args.price); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProduto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAdmin(t *testing.T) {
	tests := []struct {
		name string
		want *User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAdmin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAdmin() = %v, want %v", got, tt.want)
			}
		})
	}
}
