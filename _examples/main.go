package main

import (
	"fmt"
	"github.com/fahadalisarwar1/WebScrapingCourse/utils"
	"io/ioutil"
	"net/http"
)

func main(){
	baseURL := "https://www.youtube.com/channel/UCDY_0WzkHyj0A1ev6RTql1Q"
	//baseURL := "https://youtube.com/jsfunc"


	res, err := http.Get(baseURL)

	utils.DisplayError(err)
	body, err := ioutil.ReadAll(res.Body)
	utils.DisplayError(err)
	fmt.Println(len(body))
	fmt.Println(string(body))
	res.Body.Close()
}
