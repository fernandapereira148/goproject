package main

import (
	"br.com.industrial/loja/routes"
	"net/http"
)

func main() {
	routes.CarregarRotas()
	http.ListenAndServe(":8080", nil)
}
