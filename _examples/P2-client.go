package main

import (
	"encoding/csv"
	"github.com/gocolly/colly"
	"log"
	"strings"
	"os"
	"strconv"
	"fmt"
)

func main(){

	url := "www.czone.com.pk"
	// Create a Collector specifically for Shopify
	c := colly.NewCollector(colly.AllowedDomains(url))
	fName := "data.csv"
	// create a file
	file, err := os.Create(fName)
	// check for errors
	if err != nil {
		log.Fatalf("Could not create file, error : %q", err)
		return
	}
	// close file afterwards
	defer file.Close()

	writer := csv.NewWriter(file)
	// flush contents afterwards
	defer writer.Flush()


	c.OnHTML(".product", func(e *colly.HTMLElement) {
		links := e.ChildTexts("a")

		prices := e.ChildTexts("span")
		var camprice string
		for _, price := range prices{
			if strings.HasPrefix(price, "Rs"){
				camprice = price
			}

		}
		fmt.Println(links[1], "\t", camprice)

		writer.Write([]string{links[1], camprice})
	})



	baseURL := "https://www.czone.com.pk/cameras-pakistan-ppt.136.aspx?page="
	for i := 1; i < 20; i++ {
		//fmt.Printf("Scraping Page : %d\n", i)
		// visit each page and scrape data
		c.Visit(baseURL + strconv.Itoa(i))
	}
	log.Print("scraping complete")
	log.Println(c)
}
