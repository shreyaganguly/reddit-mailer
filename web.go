package main

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"strings"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	RootTemplate.Execute(w, nil)
}

func InsertHandler(w http.ResponseWriter, r *http.Request) {
	var newSubscriber mail.Address
	newSubscriber.Name = r.FormValue("Name")
	newSubscriber.Address = r.FormValue("Email-id")
	if newSubscriber.Address == "" {
		fmt.Fprintf(w, "Sorry the email field cannot be left empty")
	} else if isPresent(newSubscriber.Address) == true {
		fmt.Fprintf(w, "Sorry the email address is already taken")
	} else {
		AddSubscriber(newSubscriber)
		fmt.Fprintf(w, "You have succesfully subscribed")
	}
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have succesfully unsubscribed")
	unsub := strings.TrimPrefix(r.URL.String(), "/unsubscribe/")
	FindAndRemove(unsub)
}

func WebServer(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)

	http.HandleFunc("/subscribe", InsertHandler)
	http.HandleFunc("/unsubscribe/", DeleteHandler)
	http.HandleFunc("/", rootHandler)

	log.Println("Starting web interface at", addr)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatal("ERROR:", err)
	}
}
