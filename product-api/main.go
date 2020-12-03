package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/products", SetMiddlewareJSON(productHandler)).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("svc")
	JSON(w, http.StatusOK, url)
}

type Product struct {
	Id       int
	Name     string
	Comments *[]Comment
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	p := Product{
		Id:   1,
		Name: "Çanta",
	}
	p2 := Product{
		Id:   2,
		Name: "Kitap",
	}

	comments, err := GetComments()

	if err != nil {
		JSON(w, http.StatusOK, err.Error())
	}

	p.Comments = comments

	products := []*Product{&p, &p2}

	JSON(w, http.StatusOK, products)
}

type Comment struct {
	Id        int64
	ProductId int64
	Content   string
}

func GetComments() (*[]Comment, error) {
	var err error
	var comments []Comment
	url := os.Getenv("svc")
	res, err := http.Get(url + "/comments")

	if err != nil {
		return &comments, err
	}
	data, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()

	json.Unmarshal([]byte(data), &comments)

	return &comments, nil
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}
