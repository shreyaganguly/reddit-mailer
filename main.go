package main

import (
	"flag"
	"log"
	"net/mail"
	"time"
)



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
func FetchAndDispatch(mailer FeedMailer, addr []mail.Address) {
	payloadnew, err := GetFeed("golang")
	if err != nil {
		log.Println("Error encountered")
	}
	if len(addr) > 0 {
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
		go FetchAndDispatch(mailer, subscriptionList)
	}
}

var (
	server     = flag.String("s", "", "Host name of the SMTP Server")
	port       = flag.Int("t", 0, "SMTP port")
	auth_user  = flag.String("u", "", "Username for SMTP authentication")
	password   = flag.String("p", "", "Password for SMTP authentication")
	senderName = flag.String("sendername", "Reddit Mailer", "Sender name")
	senderMail = flag.String("sendermail", "no-reply@example.com", "Sender email")
	host       = flag.String("b", "0.0.0.0", "Binds to the specified IP")
	listenport = flag.Int("l", 8080, "Listens on the specified port")
)

func main() {
	log.Println("Starting reddit mailer")
	flag.Parse()
	mailer := NewMailer(
		*server,
		*port,
		mail.Address{*senderName, *senderMail},
		*auth_user,
		*password,
	)
	go FeedDispatcher(mailer)

	WebServer(*host, *listenport)

}
