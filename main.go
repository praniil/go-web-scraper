package main

import(
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"log"
	"encoding/csv"
)

type Book struct{
	Title string
	Price string
}

func main(){
	file, err := os.Create("file.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	headers := []string{"Title", "Price"}
	writer.Write(headers)

	fmt.Println("Hello World!")
	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com"),
	)
	c.OnRequest(func(r *colly.Request){
		fmt.Println("visiting: ", r.URL.String())
	})

	c.OnHTML(".next > a", func(e *colly.HTMLElement){
		newPage := e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(newPage)
	})

	c.OnHTML(".product_pod", func(e *colly.HTMLElement) {
		book := Book{}
		book.Title = e.ChildAttr(".image_container img", "alt")
		book.Price = e.ChildText(".price_color")
		row := []string{book.Title, book.Price}
		writer.Write(row)
	})

	startUrl := "http://books.toscrape.com"
	c.Visit(startUrl)
}