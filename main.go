package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

type NetworkInterface struct {
	Name     string `json:"name"`
	PublicIp string `json:"public_ip"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/interface/{interfaceName}", GetNetworkInterface)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a network interface API. Or, is it?!")
}

func GetNetworkInterface(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	interfaceName := vars["interfaceName"]

	networkInterface := NetworkInterface{
		Name:     interfaceName,
		PublicIp: "8.8.8.8",
	}

	json.NewEncoder(w).Encode(networkInterface)
}
