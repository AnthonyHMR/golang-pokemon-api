package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"pokemon-api/database"

	"github.com/gorilla/mux"
)

func getAllPokemons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.PokemonDb)
}

func newPokemon(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var pokemon database.Pokemon
	json.Unmarshal(reqBody, &pokemon)

	database.PokemonDb = append(database.PokemonDb, pokemon)

	json.NewEncoder(w).Encode(pokemon)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(commonMiddleware)
	myRouter.HandleFunc("/pokemons", getAllPokemons).Methods("GET")

	myRouter.HandleFunc("/pokemons", newPokemon).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("Pokemon Rest API")
	handleRequests()
}
