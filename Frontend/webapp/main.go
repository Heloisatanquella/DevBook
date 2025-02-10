package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	fmt.Println("Rodando WebApp!")
	utils.CarregarTemplates()

	r := router.Gerar()

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./webapp/assets"))))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
