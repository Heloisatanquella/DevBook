package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"webapp/src/config"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Carregar()
	fmt.Printf("Rodando WebApp na porta %d\n", config.Porta)
	utils.CarregarTemplates()

	r := router.Gerar()

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./webapp/assets"))))

	Porta := ":" + strconv.Itoa(config.Porta)

	server := &http.Server{
		Addr:         Porta,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
