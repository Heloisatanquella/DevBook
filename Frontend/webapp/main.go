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
	log.Fatal(http.ListenAndServe(":3000", r))
}
