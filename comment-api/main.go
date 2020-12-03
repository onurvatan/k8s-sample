package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/comments", SetMiddlewareJSON(commentHandler)).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8081", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, "index")
}

type Comment struct {
	Id        int
	ProductId int
	Content   string
}

func commentHandler(w http.ResponseWriter, r *http.Request) {

	c := Comment{
		Id:        1,
		ProductId: 1,
		Content:   "İyi",
	}

	c2 := Comment{
		Id:        2,
		ProductId: 1,
		Content:   "Kötü",
	}

	comments := []*Comment{&c, &c2}

	JSON(w, http.StatusOK, comments)
}

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
