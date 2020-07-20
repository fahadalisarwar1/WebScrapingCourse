package main

import (
	"fmt"
	utils2 "github.com/fahadalisarwar1/WebScrapingCourse/utils"
	"net/http"
	"time"
)



func main(){
	url := "http://example.webscraping.com/view/1"
	client := http.Client{Timeout:time.Second*30}
	res, err := client.Get(url)
	utils2.DisplayError(err)
	fmt.Println("Header")
	fmt.Println(res.Header)
	fmt.Println("Body")
	fmt.Println(res.Body)



}