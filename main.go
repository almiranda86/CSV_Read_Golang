package main

import (
	"core/main/helper"
	"core/main/service"
	"core/main/validator"
	"fmt"
	"net/http"
)

func messages() helper.SystemMessages {
	var system_messages helper.SystemMessages
	system_messages.Fill_defaults()

	return system_messages
}

func routeInvertResult(w http.ResponseWriter, r *http.Request) {
	response := service.InvertResult(r)
	fmt.Println(response)
	fmt.Fprint(w, response)
}

func routeMultiplyAllResult(w http.ResponseWriter, r *http.Request) {
	system_messages := messages()
	message, records := validator.ValidateRequest(r)

	if message != system_messages.ValidationOK {
		fmt.Fprint(w, message)
	}

	response := service.MultiplyAllResult(records)
	fmt.Println(response)
	fmt.Fprint(w, response)
}

func routeSumAllResult(w http.ResponseWriter, r *http.Request) {
	system_messages := messages()
	message, records := validator.ValidateRequest(r)

	if message != system_messages.ValidationOK {
		fmt.Fprint(w, message)
	}

	response := service.SumAllResult(records)
	fmt.Println(response)
	fmt.Fprint(w, response)
}

func routeFlattenResult(w http.ResponseWriter, r *http.Request) {
	system_messages := messages()
	message, records := validator.ValidateRequest(r)

	if message != system_messages.ValidationOK {
		fmt.Fprint(w, message)
	}

	response := service.FlattenResult(records)
	fmt.Println(response)
	fmt.Fprint(w, response)
}

func routeSimpleResult(w http.ResponseWriter, r *http.Request) {
	response := service.SimpleResult(r)
	fmt.Println(response)
	fmt.Fprint(w, response)
}

func main() {
	http.HandleFunc("/echo", routeSimpleResult)
	http.HandleFunc("/flatten", routeFlattenResult)
	http.HandleFunc("/sum", routeSumAllResult)
	http.HandleFunc("/multiply", routeMultiplyAllResult)
	http.HandleFunc("/invert", routeInvertResult)

	http.ListenAndServe(":8080", nil)
}
