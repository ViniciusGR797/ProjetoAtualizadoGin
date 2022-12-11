package webui

import (
	"encoding/json"
	"net/http"

	"product/entity"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// Funciona o front, mas não traz joga as informações no front

func RegisterUIHandlers(router *gin.Engine) {

	router.LoadHTMLGlob("./webui/dist/spa/*.html")

	router.Use(static.Serve("/webui", static.LocalFile("./webui/dist/spa", true)))
	router.Use(static.Serve("/webui/assets", static.LocalFile("./webui/dist/spa/assets", true)))
	router.Use(static.Serve("/webui/icons", static.LocalFile("./webui/dist/spa/icons", true)))

	router.GET("/webui", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}

func AuthLogin() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := entity.User{}

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"MSG": "Error to parse User from JSON", "codigo": 500}`))
			return
		}

		admin := entity.NewAdmin()
		token := entity.Token{
			Token: entity.USER_TOKEN,
		}

		if user.Username == admin.Username && user.Password == admin.Password {
			err = json.NewEncoder(w).Encode(token)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"MSG": "Error to parse Product to JSON", "codigo": 500}`))
				return
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"MSG": "Error to login, check username and password", "codigo": 401}`))
		}

	})
}
