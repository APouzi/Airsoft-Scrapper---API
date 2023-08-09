package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"golang.org/x/net/html"
)


func ReadHTML(prod chan []Product, wg *sync.WaitGroup) {
	defer wg.Done()
	request, err := http.Get("https://www.airsoftmegastore.com/Mega-Deals")
	if err != nil {
		panic(err)
	}

	htmlParser := html.NewTokenizer(request.Body)
	var startTokenDiv bool
	var readTitleTempBool bool
	lctCollection := []*Product{}
	regPrice, _ := regexp.Compile(`\$\d+(?:\.\d+)?`)

	for {
		ele := htmlParser.Next()

		if htmlParser.Err() == io.EOF {
			fmt.Println("End of File")
			break
		}

		if !startTokenDiv && ele == html.TokenType(html.StartTagToken) {
			element := htmlParser.Token()
			if len(element.Attr) < 1 {
				continue
			}
			if element.Data == "div" && element.Attr[0].Val == "pgrid" {
				startTokenDiv = true
			}
		}
		var elementProd html.Token
		if startTokenDiv {
			elementProd = htmlParser.Token()
		}
		if startTokenDiv {
			prod := Product{}
			if len(elementProd.Attr) > 1 {
				for _, v := range elementProd.Attr {
					if v.Key == "title" && v.Val != "Add to Cart" {
						// fmt.Println(v.Val)
						if strings.Contains(v.Val, "LCT") {
							readTitleTempBool = true
							prod.name = v.Val
							lctCollection = append(lctCollection, &prod)
						}
					}
				}

			}
			if elementProd.Type == html.TextToken {
				priceFound := regPrice.MatchString(elementProd.String())
				if priceFound {
					// fmt.Println(elementProd.String())
					if readTitleTempBool {
						htmlParser.Next()
						htmlParser.Next()
						Price := strings.TrimSpace(htmlParser.Token().String())
						lctCollection[len(lctCollection)-1].price = Price
						readTitleTempBool = false
					}

				}

			}
		}

	}
	for _, v := range lctCollection {
		fmt.Println(v.name, v.price)
	}
	close(prod)
	

}