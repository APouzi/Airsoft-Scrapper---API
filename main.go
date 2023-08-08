package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	request,err := http.Get("https://www.airsoftmegastore.com/Mega-Deals")
	if err != nil{
		panic(err)
	}

	htmlParser := html.NewTokenizer(request.Body)
	var startTokenDiv bool
	lctCollection := []string{}
	// checkPrice := `/(?=.)\$(([1-9][0-9]{0,2}(,[0-9]{3})*)|[0-9]+)?(\.[0-9]{1,2})?/`
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
			if elementProd.Type == html.TextToken{
				bool, _ := regexp.MatchString(`\$\d+(?:\.\d+)?`, elementProd.String())
				if bool {
					fmt.Println(elementProd.Data)
				}
				
			}
			if len(elementProd.Attr) > 1 {
				for _, v := range elementProd.Attr{
					if v.Key == "title" && v.Val != "Add to Cart"{
						// fmt.Println(v.Val)
						if strings.Contains(v.Val, "LCT"){
							lctCollection = append(lctCollection, v.Val)
							
						}
					}
				}
				
			}
		}
			
	}
	fmt.Println(lctCollection)
	
}