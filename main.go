package main

import (
	"fmt"
	"net/http"

	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/data/postgres"
)

func main() {
	db, err := postgres.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = postgres.CreateDb(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Olá, Squad 38!")
	})

	http.ListenAndServe(":8080", nil)
}
