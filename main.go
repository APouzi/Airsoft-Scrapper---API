package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	request,err := http.Get("https://www.airsoftmegastore.com/Mega-Deals")
	if err != nil{
		panic(err)
	}

	htmlParser := html.NewTokenizer(request.Body)
	var startTokenDiv bool
	for {
		ele := htmlParser.Next()

		if htmlParser.Err() == io.EOF{
			fmt.Println("done")
			break
		}
		
		if !startTokenDiv && ele == html.TokenType(html.StartTagToken){
			element := htmlParser.Token()
			if len(element.Attr) < 1{
				continue
			}
			if element.Data == "div" && element.Attr[0].Val == "pgrid"{
				startTokenDiv = true
				fmt.Println("here we go!")
			}
		}
		var elementProd html.Token
		if startTokenDiv {
			elementProd = htmlParser.Token()
		}
		if startTokenDiv{
			if len(elementProd.Attr) > 1 {
				for _, v := range elementProd.Attr{
					// fmt.Println(v)
					if v.Key == "title" && v.Val != "Add to Cart"{
						fmt.Println(v.Val)
					}
					// if v.Key == "title"{
					// 	fmt.Println(v.Val)
					// }
				}
				
			}
		}
			
	}
	
}