package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/mail"
	"strings"
)

const tplsubscriber = `
<!DOCTYPE html>
<html>
<body>
<h1>Subscriber form</h1>
<form action="/subscribe" method="POST">
  Name:<br>
  <input type="text" name="Name" value="alice johnson" >
  <br>
  Email-id:<br>
  <input type="text" name="Email-id" value="xyz@example.com">
  <br><br>
  <input type="submit" value="Subscribe">
</form>
</body>
</html>
`
func FindAndRemove(s string){
	for i,_ := range addr {
		if addr[i].Address == s{
			addr[i] = addr[len(addr)-1]
			addr = addr[:len(addr)-1]
			break
		}
	}
}

func isPresent(key mail.Address,addr []mail.Address) bool{
	flag := 0
	for _,v := range addr {
		if v.Address == key.Address{
			flag = 1
		}
}
if flag == 0{
 return false
 }else {
	 return true
 }
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("webpage").Parse(tplsubscriber)
	if err != nil {
		log.Println("ERROR!!")
	}
	t.Execute(w, nil)

}

func InsertHandler(w http.ResponseWriter, r *http.Request) {
	var newSubscriber mail.Address
	newSubscriber.Name = r.FormValue("Name")
	newSubscriber.Address = r.FormValue("Email-id")
	if newSubscriber.Address == ""{
		fmt.Fprintf(w,"Sorry the email field cannot be left empty")
	} else if isPresent(newSubscriber,addr) == true {
	fmt.Fprintf(w,"Sorry the email address is already taken")
} else {
	addr = append(addr, newSubscriber)
	fmt.Fprintf(w, "You have succesfully subscribed")
}
}
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have succesfully unsubscribed")
	unsub := strings.TrimPrefix(r.URL.String(), "/unsubscribe/")
	FindAndRemove(unsub)
}
