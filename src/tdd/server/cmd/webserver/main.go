package main

import (
	"log"
	"net/http"

	poker "tdd/server"
)

const dbName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf(`could not listen on :5000, %v`, err)
	}
}
