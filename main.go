package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprint(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprint(w, "POST Request Successful!")
	name := r.FormValue("name")
	phoneNumber := r.FormValue("phoneNumber")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprint(w, "Address = %s\n", address)
	fmt.Fprint(w, "PhoneNumber = %s\n", phoneNumber)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "hello!")
}
func main (){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)


	fmt.Printf("Starting the server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}