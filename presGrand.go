package main

import (
	//"encoding/json"
	"fmt"

	//"io/ioutil"

	"github.com/gocolly/colly"
)

func main() {
	//allGrands := make([]Grand, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("президентскиегранты.рф"),
	)

	collector.OnHTML("div.contests-directions__list > div.contest-directions-item-wrapper.contest-slider__item > ", func(element *colly.HTMLElement) {
		fmt.Print(element.Attr("href"))
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	collector.Visit("https://президентскиегранты.рф/public/contest/index")

}
