package main

import (
	"fmt"
	"net/http"

	"github.com/bagasalim/simas/api"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// if false {
	db, err := api.SetupDb()
	if err != nil {
		panic(err)
	}

	server := api.MakeServer(db)
	server.RunServer()
	// }

	// http.HandleFunc("/", test)
	// // http.HandleFunc("/ws", socket_system.WsEndpoint)
	// http.ListenAndServe(":8020", nil)
}
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	// fmt.Fprintf(w, "Home Page")
}
