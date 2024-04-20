package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strings"

// 	"github.com/PuerkitoBio/goquery"
// )

// type Sight struct {
// 	Title      string `json:"title"`
// 	Link       string `json:"link"`
// 	ButtonText string `json:"button_text"`
// }

// func fetchHTML(url string) (string, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != 200 {
// 		return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
// 	}

// 	doc, err := goquery.NewDocumentFromReader(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	var htmlContent strings.Builder
// 	doc.Find("body").Each(func(i int, s *goquery.Selection) {
// 		html, err := s.Html()
// 		if err == nil {
// 			htmlContent.WriteString(html)
// 		}
// 	})

// 	return htmlContent.String(), nil
// }

// func main() {
// 	url := "https://astana.citypass.kz/ru/category/muzei-i-galerei/"
// 	html, err := fetchHTML(url)
// 	if err != nil {
// 		log.Fatal("Failed to fetch HTML: ", err)
// 	}

// 	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
// 	if err != nil {
// 		log.Fatal("Error loading parsed HTML body. ", err)
// 	}

// 	sights := []Sight{}
// 	doc.Find(".sights__item--imfo.one_item_blok").Each(func(index int, item *goquery.Selection) {
// 		sight := Sight{
// 			Title:      item.Find(".sights__item--title a").Text(),
// 			Link:       item.Find(".sights__item--title a").AttrOr("href", ""),
// 			ButtonText: item.Find(".sights__item--img img").AttrOr("src", ""),
// 		}
// 		sights = append(sights, sight)
// 	})

// 	jsonData, err := json.MarshalIndent(sights, "", "    ")
// 	if err != nil {
// 		log.Fatal("Error marshalling data to JSON: ", err)
// 	}

// 	// Write JSON data to a file
// 	err = os.WriteFile("sights_data.txt", jsonData, 0644)
// 	if err != nil {
// 		log.Fatal("Error writing JSON to file: ", err)
// 	}

// 	fmt.Println("Data written to sights_data.txt")
// }
