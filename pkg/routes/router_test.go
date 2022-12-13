package routes

import (
	"product/pkg/service"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestConfigRoutes(t *testing.T) {
	type args struct {
		router  *gin.Engine
		service service.ProdutoServiceInterface
	}
	tests := []struct {
		name string
		args args
		want *gin.Engine
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConfigRoutes(tt.args.router, tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfigRoutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
