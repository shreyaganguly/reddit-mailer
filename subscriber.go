package main

import (
	"net/mail"
	"sync"
)

var (
	lock sync.Mutex
 	subscriptionList = make([]mail.Address, 0)
)

func AddSubscriber(s mail.Address)  {
	lock.Lock()
	subscriptionList = append(subscriptionList, s)
	lock.Unlock()
}


func FindAndRemove(s string) {
	lock.Lock()
	for i, _ := range subscriptionList {
		if subscriptionList[i].Address == s {
			subscriptionList[i] = subscriptionList[len(subscriptionList)-1]
			subscriptionList = subscriptionList[:len(subscriptionList)-1]
			break
		}
	}
	lock.Unlock()
}

func isPresent(key mail.Address, subscriptionList []mail.Address) bool {
	flag := 0
	for _, v := range subscriptionList {
		if v.Address == key.Address {
			flag = 1
		}
	}
	if flag == 0 {
		return false
	} else {
		return true
	}
}
