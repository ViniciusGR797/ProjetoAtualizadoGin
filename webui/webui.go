package webui

import (
	"encoding/json"
	"net/http"

	"product/pkg/entity"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// Função que cria as rotas com o front end (webui)
func RegisterUIHandlers(router *gin.Engine) {
	// Rota do HTML global
	router.LoadHTMLGlob("./webui/dist/spa/*.html")
	// Demais rotas e onde localiza o arquivo
	router.Use(static.Serve("/webui", static.LocalFile("./webui/dist/spa", true)))
	router.Use(static.Serve("/webui/assets", static.LocalFile("./webui/dist/spa/assets", true)))
	router.Use(static.Serve("/webui/icons", static.LocalFile("./webui/dist/spa/icons", true)))

	// Pega rota com method GET, chama o html do arquivo index.html
	router.GET("/webui", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}

// Função de autenticação de login
func AuthLogin() http.Handler {
	// retorna requesição HTTP
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// cria variável do tipo User (vazia)
		user := entity.User{}

		// converter json no user (para fazer as verificações de acesso)
		err := json.NewDecoder(r.Body).Decode(&user)
		// verifica se teve erro ao converter
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"MSG": "Error to parse User from JSON", "codigo": 500}`))
			return
		}

		// cria variável admin com os dados padrão de admin com método NewAdmin (construtor)
		admin := entity.NewAdmin()

		// cria variável token do tipo Token, recebe o token fake
		token := entity.Token{
			Token: entity.USER_TOKEN,
		}

		// verifica se o username e a senha estão corretas
		if user.Username == admin.Username && user.Password == admin.Password {
			// converter json no token
			err = json.NewEncoder(w).Encode(token)
			// verifica se teve erro ao converter
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"MSG": "Error to parse Product to JSON", "codigo": 500}`))
				return
			}
		} else {
			// erro de credenciais (username ou senha incorreta)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"MSG": "Error to login, check username and password", "codigo": 401}`))
		}

	})
}
