package main

import (
	"reflect"
	"testing"

	// Import interno de packages do próprio sistema
	"product/pkg/entity"
	"product/pkg/service"
)

// Teste unitário do método main
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

// Teste unitário do método ConnectBD
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

// Teste unitário do método CreateExcel
func TestCreateExcel(t *testing.T) {
	type args struct {
		list_log *entity.LogList
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateExcel(tt.args.list_log)
		})
	}
}
