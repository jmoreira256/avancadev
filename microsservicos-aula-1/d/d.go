package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Status string
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	provaMS4 := r.PostFormValue("provaMS4")
	result := returnResultado(provaMS4)
	fmt.Print(result)
}

func returnResultado(provaMS4 string) Result {

	result := Result{}

	return result
}
