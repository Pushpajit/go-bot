package scrapping

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Image struct {
	Title string
	URL   string
}

func GetImage(msg []string) []Image {
	var images []Image

	images = append(images, getImageDesktopHut(msg)...)
	images = append(images, getImageUnsplash(msg)...)

	return images
}

func getImageUnsplash(msg []string) []Image {
	var images []Image

	// Start the performance timer
	startTime := time.Now()

	c := colly.NewCollector(
		colly.AllowedDomains("unsplash.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Scrapping the URL: %v\n", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Fatal(err.Error())
	})

	c.OnHTML("div.wdUrX img[class='I7OuT DVW3V L1BOa']", func(h *colly.HTMLElement) {
		// Start sending the scrapped URL into the channel, then it will also start downloading.
		images = append(images, Image{
			Title: h.Attr("alt"),
			URL:   h.Attr("src"),
		})
	})

	URL := fmt.Sprintf("https://unsplash.com/s/photos/%v?license=free&orientation=landscape", msg[1])

	// check for the orientation
	if len(msg) == 4 && strings.ToLower(msg[3]) == "portrait" {
		URL = fmt.Sprintf("https://unsplash.com/s/photos/%v?license=free&orientation=portrait", msg[1])
	}

	c.Visit(URL) // Initiate the web sracpping

	elapsed := time.Since(startTime)
	fmt.Printf("💿 All Download Completed ✅\n⏳ Time Taken %vs \n", elapsed.Seconds())

	return images
}

func getImageDesktopHut(msg []string) []Image {
	var images []Image

	c := colly.NewCollector(
		colly.AllowedDomains("www.desktophut.com"),
	)

	// It will trigger when it first hit the URL.
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Scrapping the URL: %v\n", r.URL)
	})

	// To handle the errors during scrapping.
	c.OnError(func(r *colly.Response, err error) {
		log.Fatal(err.Error())
	})

	// Start sending the scrapped URL into the channel, then it will also start downloading.
	c.OnHTML("img[class='customimg img-fluid rounded  ']", func(h *colly.HTMLElement) {
		images = append(images, Image{
			Title: h.Attr("alt"),
			URL:   h.Attr("src"),
		})
	})

	// Format the URL
	URL := fmt.Sprintf("https://www.desktophut.com/search/%v", strings.Join(strings.Split(msg[1], "-"), " "))

	// Start visiting the URL
	c.Visit(URL)

	return images

}
