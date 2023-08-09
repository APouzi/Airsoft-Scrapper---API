package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
)

func Ebay() {
	url := "https://api.ebay.com/buy/browse/v1/item/?"
	url += "seller=airsoftmegastorecom"
	response, err := http.Get(url)
	if err != nil{
		fmt.Println("Ebay issue")
		panic(err)
	}
	
	bReader := bufio.NewReader(response.Body)
	byteInsert := make([]byte, 500)
	for {
		n, err := bReader.Read(byteInsert)
		
		if err == io.EOF{
			fmt.Println("done")
			break
		}
		fmt.Println(string(byteInsert[:n]))
	}

}