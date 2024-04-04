package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type singleTon struct{}

var singleTonInstance *singleTon

func GetInstance() *singleTon {
	if singleTonInstance == nil {
		// accuire lock to create an instance
		lock.Lock()
		defer lock.Unlock()

		if singleTonInstance == nil {
			fmt.Println("Creating a SingleTon Instance!")
			singleTonInstance = &singleTon{}
		} else {
			fmt.Println("Instance is already created!!!")
		}
	} else {
		fmt.Println("Instance is already created!!!")
	}

	return singleTonInstance
}
