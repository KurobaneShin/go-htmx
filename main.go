package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)
type Test struct {
	Title  string
	Description string
}
func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		templates := []string{
			"static/header.html",
			"static/index.html",
			"static/footer.html",
		}

		t := template.Must(template.New("index.html").ParseFiles(templates...))

		data := map[string][]Test {
			"Data":{
				{Title: "Title 1", Description: "Description 1"},
			},
		}

		err := t.Execute(w, data)

		if err != nil {
			log.Fatal(err)
		}
	})

	mux.HandleFunc("/post", handlePost)

	log.Fatal(http.ListenAndServe(":8080", mux))
}



func handlePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var data = Test{}

	decoder := schema.NewDecoder()
	err := decoder.Decode(&data, r.Form)
	if err != nil {
		panic(err)
	}


	t := template.Must(template.ParseFiles("static/index.html"))
	t.ExecuteTemplate(w,"list-element", data)

	err2 := t.Execute(w, data)
	if err != nil {
		log.Fatal(err2)
	}
}
