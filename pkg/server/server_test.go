package server

import (
	"product/config"
	"product/pkg/service"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

// Teste unitário do método NewServer
func TestNewServer(t *testing.T) {
	type args struct {
		conf *config.Config
	}
	tests := []struct {
		name string
		args args
		want Server
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Teste unitário do método Run
func TestRun(t *testing.T) {
	type args struct {
		router  *gin.Engine
		server  Server
		service service.ProdutoServiceInterface
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Run(tt.args.router, tt.args.server, tt.args.service)
		})
	}
}
