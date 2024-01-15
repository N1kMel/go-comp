package main

import (
	"encoding/json"

	"fmt"
	"log"

	"os"

	//"io/ioutil"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Grand struct {
	//ID        int    `json:"id"`
	Name      string    `json:"name"`
	DateStart time.Time `json:"date_start"`
	DateEnd   time.Time `json:"date_end"`
	Organizer string    `json:"organizer"`
	//URL       string `json:"URL"`
}

func extractDates(input string) (time.Time, time.Time, error) {
	parts := strings.Fields(input)

	startDateStr := parts[1] + " " + parts[2] + " " + time.Now().Format("2006")
	endDateStr := parts[4] + " " + parts[5] + " " + time.Now().Format("2006")

	startDate, err1 := time.Parse("2 Jan 2006", startDateStr)
	endDate, err2 := time.Parse("2 Jan 2006", endDateStr)

	if err1 != nil || err2 != nil {
		x := fmt.Errorf("Error parsing dates: %v, %v", err1, err2)
		return time.Time{}, time.Time{}, x
	}

	return startDate, endDate, nil
}

func main() {

	org := "Фонд президентских грантов"

	collector := colly.NewCollector(
		colly.AllowedDomains("президентскиегранты.рф"),
	)

	dateInf := make([]string, 0)
	collector.OnHTML("div.acceptance-projects-competitions__item_info > p", func(h *colly.HTMLElement) {
		dateInf = append(dateInf, h.Text)
		fmt.Print(dateInf)
	})

	var dateStart time.Time
	var dateEnd time.Time

	dateStart, err1 := time.Parse("2006-Jan-02", "2024-Feb-01")
	dateEnd, err2 := time.Parse("2006-Jan-02", "2024-Mar-15")
	if err1 != nil || err2 != nil {
		log.Fatal(err1, err2)
	}

	allGrands := make([]Grand, 0)

	collector.OnHTML("div.contest-directions-item-wrapper.contest-slider__item a", func(h *colly.HTMLElement) {
		//fmt.Print(h.Attr("href"))
		collector.Visit(h.Request.AbsoluteURL(h.Attr("href")))
	})
	collector.OnHTML("ol.contest-subdir__list.green-styled > li", func(element *colly.HTMLElement) {
		//fmt.Print(element.Text)
		grand := Grand{
			//ID: GrandId,
			Name:      element.Text,
			DateStart: dateStart,
			DateEnd:   dateEnd,
			Organizer: org,
		}
		allGrands = append(allGrands, grand)
	})

	collector.OnRequest(func(request *colly.Request) {
		//fmt.Println("Visiting", request.URL.String())
	})

	collector.Visit("https://президентскиегранты.рф/public/contest/index")

	//fmt.Println(allGrands)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(allGrands)
}
