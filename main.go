package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("index.html").ParseFiles("static/index.html"))

		data := map[string]interface{}{"Test": "Hello World", "Test2": "Hello World 2"}
		err := t.Execute(w, data)
		if err != nil {
			log.Fatal(err)
		}
	})

	mux.HandleFunc("/post", handlePost)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

type Test struct {
	Test  string
	Test2 string
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println("method", r.Method)

	var data = Test{}

	decoder := schema.NewDecoder()

	err := decoder.Decode(&data, r.Form)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)

	t := template.Must(template.New("index.html").ParseFiles("static/index.html"))

	err2 := t.Execute(w, data)
	if err != nil {
		log.Fatal(err2)
	}
}
