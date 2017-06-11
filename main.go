package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.POST("/register", register)
	// router.POST("/register", register())

	log.Fatal(http.ListenAndServe(":8888", router))
}

func index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(r.Form["title"])
	// fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // write data to response
}
