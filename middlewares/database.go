package middlewares

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/iquark/go-api-sample/models"
)

// DB definition of DB type
type DB struct {
	models.DataStore
}

func DatabaseMiddleware(next http.Handler) http.Handler {
	dbInstance, err := models.NewDB()

	if err != nil {
		log.Panic(err)
	} else {
		log.Println("BD ok")
	}

	db := &DB{dbInstance}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, "DB", db)
		next.ServeHTTP(w, r)
	})
}
