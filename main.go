package main

import (
	"fmt"
	"sync"
)

type Product struct {
	name  string
	price string
}

func main() {
	listCh := make(chan []Product)
	var wg sync.WaitGroup
	wg.Add(1)
	go ReadHTML(listCh, &wg)
	wg.Wait()
	for _, v := range <-listCh {
		fmt.Println(v.name, v.price)
	}
	// Ebay()
}
