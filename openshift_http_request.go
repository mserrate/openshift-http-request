package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
	response := os.Getenv("RESPONSE")
	
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
	  fmt.Println(err)
	}
	
	if len(response) == 0 {
		response = "Openshift Http Test"
	} else {
		response = string(requestDump)
	}

	fmt.Fprintln(w, response)
	fmt.Println("Servicing request.")

	fmt.Println(string(requestDump))
}

func listenAndServe(port string) {
	fmt.Printf("serving on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func main() {
	http.HandleFunc("/", httpHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	go listenAndServe(port)

	port = os.Getenv("SECOND_PORT")
	if len(port) == 0 {
		port = "8888"
	}
	go listenAndServe(port)

	select {}
}