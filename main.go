package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"bytes"

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

	publicIpAddress := GetNetworkInterfacePublicIp(interfaceName)

	networkInterface := NetworkInterface{
		Name:     interfaceName,
		PublicIp: publicIpAddress,
	}

	json.NewEncoder(w).Encode(networkInterface)
}

func Ifup(interfaceName string) {
        exec.Command("sudo", "ifup", interfaceName).Run()
}

func Ifdown(interfaceName string) {
        exec.Command("sudo", "ifdown", interfaceName).Run()
}

func GetNetworkInterfacePublicIp(interfaceName string) string {
	var (
		cmdOut []byte
		err    error
		ipifyResult struct {
			Ip  string         `json:"ip"`
		}
	)

	if cmdOut, err = exec.Command("curl", "--interface", interfaceName, "https://api.ipify.org?format=json").Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running curl --interface command: ", err)
	}
	json.NewDecoder(bytes.NewBuffer(cmdOut)).Decode(&ipifyResult)

	return ipifyResult.Ip
}
