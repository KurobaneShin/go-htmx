package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/KurobaneShin/go-htmx.git/database"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type Test struct {
	Title       string
	Description string
}

func main() {

	mux := mux.NewRouter().StrictSlash(true)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		templates := []string{
			"static/header.html",
			"static/index.html",
			"static/footer.html",
		}

		t := template.Must(template.New("index.html").ParseFiles(templates...))

		list := database.GetList()

		data := map[string][]database.ListItem{
			"Data": list,
		}

		err := t.Execute(w, data)

		if err != nil {
			log.Fatal(err)
		}
	})

	mux.HandleFunc("/post", handlePost).Methods("Post")

	mux.HandleFunc("/put/{id}", handlePutAction).Methods("Put")
	mux.HandleFunc("/put/{id}", handlePutData)

	mux.HandleFunc("/delete/{id}", handleDelete).Methods("Delete")

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

	listItem := database.InsertListItem(data.Title, &data.Description)

	t := template.Must(template.ParseFiles("static/index.html"))
	t.ExecuteTemplate(w, "list-element", listItem)

	err2 := t.Execute(w, data)
	if err != nil {
		log.Fatal(err2)
	}
}

func handlePutData(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("static/edit.html"))

	id := mux.Vars(r)["id"]

	castedId, castErr := strconv.ParseInt(id, 10, 64)

	if castErr != nil {
		panic(castErr)
	}

	item := database.ReadListItem(castedId)

	err := t.Execute(w, item)
	if err != nil {
		log.Fatal(err)
	}
}

func handlePutAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var data = Test{}

	decoder := schema.NewDecoder()
	err := decoder.Decode(&data, r.Form)

	if err != nil {
		panic(err)
	}

	id := mux.Vars(r)["id"]

	castedId, castErr := strconv.ParseInt(id, 10, 64)

	if castErr != nil {
		panic(castErr)
	}

	database.UpdateListItem(castedId, data.Title, &data.Description)

	listItem := database.ListItem{Id: castedId, Title: data.Title, Description: &data.Description}

	t := template.Must(template.ParseFiles("static/index.html"))
	t.ExecuteTemplate(w, "list-element", listItem)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	castedId, castErr := strconv.ParseInt(id, 10, 64)

	if castErr != nil {
		panic(castErr)
	}

	database.DeleteListItem(castedId)
}
