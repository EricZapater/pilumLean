package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Profile struct {
	Name    string
	Hobbies []string
}

type Kpi struct {
	Descripcio string
	Objectiu   float32
	Valor      float32
}



func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/kpi", kpiHander)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func kpiHander(w http.ResponseWriter, r *http.Request) {
	kpi := []Kpi{
		Kpi{"OEE", 88.0, 88.1},
		Kpi{"Prod", 95, 88},
	}
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	jskpi, err := json.Marshal(kpi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jskpi)
}
