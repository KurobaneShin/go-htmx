package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

type ListItem struct {
	Id          int64
	Title       string
	Description *string
}

func GetConnection() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err.Error())
	}

	return db, nil
}

func GetList() []ListItem {
	db, connErr := GetConnection()
	CheckError(connErr)

	rows, err := db.Query(`SELECT "id", "title", "description" from  "list"`)
	CheckError(err)
	defer rows.Close()

	var list []ListItem

	for rows.Next() {
		var item ListItem
		err := rows.Scan(&item.Id, &item.Title, &item.Description)

		CheckError(err)

		list = append(list, item)
	}

	return list
}

func InsertListItem(title string, description *string) ListItem {
	db, connErr := GetConnection()
	CheckError(connErr)

	stmt, err := db.Prepare("INSERT INTO list (title, description) VALUES ($1, $2)")

	CheckError(err)

	res, e := stmt.Exec(title, description)

	CheckError(e)

	id, _ := res.LastInsertId()

	return ListItem{Id: id, Title: title, Description: description}

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
