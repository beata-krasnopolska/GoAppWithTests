package main

import (
	"log"
	"net/http"
	"poker"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.NewFileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	defer close()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
	// if err := http.ListenAndServe(":5000", server); err != nil {
	// 	log.Fatalf("could not listen on port 5000 %v", err)
	// }

	// server := NewPlayerServer(NewInMemoryPlayerStore())
	// log.Fatal(http.ListenAndServe(":5000", server))

	// flag.Parse()
	// if *clientID == "" || *clientSecret == "" {
	// 	fmt.Println("You must specify the client ID and client secret from https://api.slack.com/applications")
	// 	os.Exit(1)
	// }
	// http.HandleFunc("/add", addToSlack)
	// http.HandleFunc("/auth", auth)
	// http.HandleFunc("/", home)
	// log.Fatal(http.ListenAndServe(*address, nil))
}
