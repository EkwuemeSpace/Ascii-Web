package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)

	//Look inside the folder named static in my project.
	fs := http.FileServer(http.Dir("./static"))
	//Turn my folder into a mini web server for files
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("staring seerver on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
		return
	}
}
