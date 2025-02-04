package main

import ( // Adjust based on your module path
	"log"
	"net/http"
)

func main() {
	setupAPI()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupAPI() {
	manager := NewManager() // This should now work

	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.serveWS)
}
