package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func recipesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		recipes, _, _ := getValidRecipes(false)
		// Just send out the JSON version of 'tom'
		j, _ := json.Marshal(recipes)
		enableCors(&w)
		w.Write(j)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "I can't do that.")
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func server() {
	http.HandleFunc("/recipes", recipesHandler)
	log.Println("and party and bullshit and")
	http.ListenAndServe(":8080", nil)
}
