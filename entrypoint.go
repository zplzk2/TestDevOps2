package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	port = flag.Uint("port", 50721, "The port to listen on")
)

func printBanner() {
	fmt.Println("================================================================")
	fmt.Println("= TestDevOps2")
	fmt.Println("= Port", *port)
	fmt.Println("================================================================")
}

func add(w http.ResponseWriter, r *http.Request) {
	num1, _ := strconv.Atoi(mux.Vars(r)["num1"])
	num2, _ := strconv.Atoi(mux.Vars(r)["num2"])

	result := num1 + num2

	json.NewEncoder(w).Encode(result)

	log.Printf("add %v and %v, got %v", num1, num2, result)
}

func main() {
	flag.Parse()

	printBanner()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/add/{num1}/{num2}", add).Methods("GET")
	addr := fmt.Sprintf(":%v", *port)

	log.Fatal(http.ListenAndServe(addr, router))
}
