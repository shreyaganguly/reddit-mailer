package main

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"time"
	"flag"
)

var addr = []mail.Address{
	mail.Address{"Divya Surana", "divyas@rssoftware.co.in"},
	mail.Address{"Surabhi Gupta", "surabhig@rssoftware.co.in"},
	mail.Address{"Khusbu Mishra", "khusbum@rssoftware.co.in"},
	mail.Address{"Shreya Ganguly", "shreya.ganguly99@gmail.com"},
	mail.Address{"Anirban Gupta", "anirbang@rssoftware.co.in"},
	mail.Address{"Pratik Kumar Singh", "pratiks@rssoftware.co.in"},
}
var payloadold Feed


func FindDifference(feedold Feed, feednew []Feed) int {
	var index int
	for i, _ := range feednew {
		if feednew[i].URL == feedold.URL && feednew[i].Title == feedold.Title {
			index = i
		}
	}

	if index == 0 {
		return -1
	}
	return index

}
func FetchAndDispatch(mailer FeedMailer) {
	payloadnew, err := GetFeed("golang")
	if err != nil {
		log.Println("Error encountered")
	}
	if len(payloadold.URL) == 0 && len(payloadold.Title) == 0 {
		mailer.Send(addr, payloadnew)
	}
	index := FindDifference(payloadold, payloadnew)
	payloadold.URL = payloadnew[0].URL
	payloadold.Title = payloadnew[0].Title
	if index != -1 {
		mailer.Send(addr, payloadnew)
	}
}

func FeedDispatcher(mailer FeedMailer) {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	for range ticker.C {
		go FetchAndDispatch(mailer)
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", *r)
}

func main() {
	log.Println("Starting reddit mailer")
	var server = flag.String("server", "", "for giving the server name")
	var port  = flag.Int("port",0,"for giving the port number")
	var password = flag.String("password","XXXX","for giving the password")
	var auth_user = flag.String("auth-user","","for giving the authorized user email")
	var senderMail= flag.String("sendermail", "mail", "sender email")
	var senderName= flag.String("sendername", "name", "sender name")
	flag.Parse()
	mailer := NewMailer(
		*server,
		*port,
		mail.Address{*senderName, *senderMail},
		*auth_user,
		*password,
	)
	go FeedDispatcher(mailer)

	http.HandleFunc("/", RootHandler)
	http.ListenAndServe(":8080", nil)
}
