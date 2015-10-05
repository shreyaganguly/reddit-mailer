package main

import (
	"net/mail"
	"sync"
)

var (
	lock sync.Mutex
 	subscriptionList = make(map[string]mail.Address)
)

func AddSubscriber(s mail.Address)  {
	lock.Lock()
	subscriptionList[s.Address] = s
	lock.Unlock()
}


func FindAndRemove(s string) {
	lock.Lock()
	delete(subscriptionList,s)
	lock.Unlock()
}

func isPresent(key string) bool {
	_,ok :=subscriptionList[key]
	return ok
}
