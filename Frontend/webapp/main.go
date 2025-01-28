package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	fmt.Println("Rodando WebApp!")
	utils.CarregarTemplates()

	r := router.Gerar()

	// Servir os arquivos estáticos corretamente
	fileServer := http.FileServer(http.Dir("./webapp/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	log.Fatal(http.ListenAndServe(":3000", r))
}
