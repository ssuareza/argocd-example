package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		env := os.Getenv("ENVIRONMENT")
		if env == "" {
			env = "ENVIRONMENT variable not set"
		}
		fmt.Fprintf(w, "Hello from argocd-example! Reading variable from %s\n", env)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
