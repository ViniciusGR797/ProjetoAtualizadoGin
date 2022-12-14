package main

import (
	"product/pkg/entity"
	"product/pkg/service"
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestConnectBD(t *testing.T) {
	tests := []struct {
		name string
		want *service.Produto_service
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConnectBD(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectBD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateExcel(t *testing.T) {
	type args struct {
		list_product *entity.ProdutoList
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateExcel(tt.args.list_product)
		})
	}
}
