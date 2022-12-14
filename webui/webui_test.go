package webui

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

// Teste unitário do método RegisterUIHandlers
func TestRegisterUIHandlers(t *testing.T) {
	type args struct {
		router *gin.Engine
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterUIHandlers(tt.args.router)
		})
	}
}

// Teste unitário do método AuthLogin
func TestAuthLogin(t *testing.T) {
	tests := []struct {
		name string
		want http.Handler
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AuthLogin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}
