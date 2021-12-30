package routes

import (
	"go-rest-api/controllers"
	"go-rest-api/middlewares"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Init() {
	router := mux.NewRouter()
	router.Use(middlewares.ContentTypeMiddleware)
	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/personalities", controllers.List).Methods("Get")
	router.HandleFunc("/personalities/{id}", controllers.Get).Methods("Get")
	router.HandleFunc("/personalities", controllers.Create).Methods("Post")
	router.HandleFunc("/personalities/{id}", controllers.Delete).Methods("Delete")
	router.HandleFunc("/personalities/{id}", controllers.Update).Methods("Put")

	http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router))
}
