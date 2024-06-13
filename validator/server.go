package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func cuisinesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		_, cuisines, _ := getValidRecipes(false)
		// Just send out the JSON version of 'tom'
		j, err := json.Marshal(cuisines)
		if err != nil {
			fmt.Println("davo big problem looky")
		}
		enableCors(&w)
		w.Write(j)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "I can't do that.")
	}
}

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
	http.HandleFunc("/cuisines", cuisinesHandler)
	log.Println("and party and bullshit and")
	http.ListenAndServe(":8080", nil)
}
