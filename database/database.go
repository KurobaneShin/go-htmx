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
Id int64
  Title string
  Description *string
}

func GetConnection() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

  db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err.Error())
	}


    // insert
    // hardcoded
    // insertStmt := `insert into "Students"("Name", "Roll") values('John', 1)`
    // _, e := db.Exec(insertStmt)
    // CheckError(e)

	return db, nil
}

func GetList() []ListItem {
  db,connErr := GetConnection()
  CheckError(connErr)

  rows,err:= db.Query(`SELECT "id", "title", "description" from  "list"`)
  CheckError(err)
  defer rows.Close()

  var list[]ListItem

  for rows.Next() {
    var item ListItem
    err:= rows.Scan(&item.Id,&item.Title,&item.Description)

    CheckError(err)

    list = append(list, item)
  }

  return list

}

func CheckError(err error){
  if err != nil {
    panic(err)
  }
}
