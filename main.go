package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

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
    fmt.Fprintln(w, "Interface name:", interfaceName)
}
