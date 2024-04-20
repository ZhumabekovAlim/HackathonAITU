package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type SightDetail struct {
	Name         string `json:"name"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phone_number"`
	ContentInfo  string `json:"content_info"`
	BusNumbers   string `json:"bus_numbers"`
	WorkingDays  string `json:"working_days"`
	WorkingHours string `json:"working_hours"`
}

type Sight struct {
	Title      string `json:"title"`
	Link       string `json:"link"`
	ButtonText string `json:"button_text"`
}

func fetchSightDetails(url string) (*SightDetail, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	detail := &SightDetail{
		Name:        doc.Find(".object__title").Text(),
		Address:     doc.Find(".object__info--adres").Text(),
		PhoneNumber: doc.Find(".object__info--email a").AttrOr("href", ""),
		ContentInfo: doc.Find(".object_content--desc").Text(),
		BusNumbers: strings.Join(doc.Find(".object_content--right-list.object_content--get span").Map(func(i int, s *goquery.Selection) string {
			return s.Text()
		}), ", "),
		WorkingDays:  doc.Find(".object_content--right-list.object_content--timetable .object_content-one").Text(),
		WorkingHours: doc.Find(".object_content--right-list.object_content--timetable .object_content-too").Text(),
	}

	return detail, nil
}

func main() {
	// Read JSON file
	fileBytes, err := ioutil.ReadFile("sights_data.txt")
	if err != nil {
		log.Fatal("Error reading JSON file: ", err)
	}

	var sights []Sight
	if err := json.Unmarshal(fileBytes, &sights); err != nil {
		log.Fatal("Error unmarshaling JSON: ", err)
	}

	var details []*SightDetail
	for _, sight := range sights {
		detail, err := fetchSightDetails(strings.TrimSpace(sight.Link))
		if err != nil {
			log.Printf("Error fetching sight details for URL %s: %v\n", sight.Link, err)
			continue
		}
		details = append(details, detail)
	}

	// Open file for writing
	file, err := os.Create("sight_details.txt")
	if err != nil {
		log.Fatal("Error creating text file: ", err)
	}
	defer file.Close()

	// Write details in plain text format
	for _, detail := range details {
		fmt.Fprintf(file, "Name: %s\n", detail.Name)
		fmt.Fprintf(file, "Address: %s\n", detail.Address)
		fmt.Fprintf(file, "Phone Number: %s\n", detail.PhoneNumber)
		fmt.Fprintf(file, "Content Information: %s\n", detail.ContentInfo)
		fmt.Fprintf(file, "Bus Numbers: %s\n", detail.BusNumbers)
		fmt.Fprintf(file, "Working Days: %s\n", detail.WorkingDays)
		fmt.Fprintf(file, "Working Hours: %s\n\n", detail.WorkingHours)
	}

	fmt.Println("Sight details saved to 'sight_details.txt'")
}
