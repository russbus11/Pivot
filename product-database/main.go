package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
)

var Item []struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	os.Remove("products.db")

	file, err := os.Create("products.db")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	db, err := sql.Open("sqlite3", "products.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table products (id integer not null primary key, name text, price real);
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	byteSlice, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(byteSlice, &Item); err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into products(id, name, price) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, product := range Item {
		_, err = stmt.Exec(product.ID, product.Name, product.Price)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i <= 5; i++ {
		var id int
		var name string
		var price float64
		rows, err := db.Query("select id, name, price from products where id = ?", i)

		for rows.Next() {
			err = rows.Scan(&id, &name, &price)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("ID: %v | Name: %v | Price: %v", id, name, price)
		}
	}
}
