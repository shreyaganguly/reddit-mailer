package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"time"
)

var addr = make([]mail.Address,0)

var payloadold Feed

func FindDifference(feedold Feed, feednew []Feed) int {
	var index int
	if len(feedold.URL) > 0 && len(feedold.Title) > 0 {
	for i, _ := range feednew {
		if feednew[i].URL == feedold.URL && feednew[i].Title == feedold.Title {
			index = i
		}
	}
}
	if index == 0 {
		return -1
	}
	return index

}
func FetchAndDispatch(mailer FeedMailer,addr []mail.Address) {
	payloadnew, err := GetFeed("golang")
	if err != nil {
		log.Println("Error encountered")
	}
	if len(addr) >0{
	index := FindDifference(payloadold, payloadnew)
	if index != -1 {
		mailer.Send(addr, payloadnew[:index])
	}
payloadold.URL = payloadnew[0].URL
payloadold.Title = payloadnew[0].Title
}
}

func FeedDispatcher(mailer FeedMailer) {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	for range ticker.C {
		go FetchAndDispatch(mailer,addr)
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", r.URL)
}

func main() {
	log.Println("Starting reddit mailer")
	var server = flag.String("server", "", "for giving the server name")
	var port = flag.Int("port", 0, "for giving the port number")
	var password = flag.String("password", "XXXX", "for giving the password")
	var auth_user = flag.String("auth-user", "", "for giving the authorized user email")
	var senderMail = flag.String("sendermail", "mail", "sender email")
	var senderName = flag.String("sendername", "name", "sender name")
	var host = flag.String("b", "0.0.0.0", "Binds to the specified IP")
  var listenport = flag.Int("p", 8080, "Listens on the specified port")
	flag.Parse()
	mailer := NewMailer(
		*server,
		*port,
		mail.Address{*senderName, *senderMail},
		*auth_user,
		*password,
	)
	go FeedDispatcher(mailer)

  addr := fmt.Sprintf("%s:%d",*host,*listenport)
	http.HandleFunc("/subscribe", InsertHandler)
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/unsubscribe/",DeleteHandler)
	http.ListenAndServe(addr, nil)
}
