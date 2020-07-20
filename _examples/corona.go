package main

import (
	"encoding/csv"
	"github.com/fahadalisarwar1/WebScrapingCourse/utils"
	"github.com/gocolly/colly"
	"log"
	"os"
)

func main(){
	url := "https://en.wikipedia.org/wiki/Foreign_trade_of_Pakistan"

	fName := "data.csv"

	file, err := os.Create(fName)
	utils.DisplayError(err)

	defer file.Close()
	writer := csv.NewWriter(file)

	c := colly.NewCollector(
		colly.AllowedDomains(url),
	)

	query := ".wikitable sortable jquery-tablesorter"
	c.OnHTML(query, func(e *colly.HTMLElement){

		res := e.ChildAttr("table", "a")

		writer.Write([]string{res,})
		//fmt.Println(e.ChildText("a"))
	})

	c.Visit(url)
	writer.Flush()
	log.Print("scraping complete")
	log.Println(c)
}
