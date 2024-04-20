package main

import (
	"encoding/csv"
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
	Name         string
	Address      string
	PhoneNumber  string
	ContentInfo  string
	BusNumbers   string
	WorkingDays  string
	WorkingHours string
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

	cleanText := func(s string) string {
		// Remove new lines and trim space
		return strings.TrimSpace(strings.ReplaceAll(s, "\n", ""))
	}

	detail := &SightDetail{
		Name:        cleanText(doc.Find(".object__title").Text()),
		Address:     cleanText(doc.Find(".object__info--adres").Text()),
		PhoneNumber: cleanText(doc.Find(".object__info--email a").AttrOr("href", "")),
		ContentInfo: cleanText(doc.Find(".object_content--desc").Text()),
		BusNumbers: strings.Join(doc.Find(".object_content--right-list.object_content--get span").Map(func(i int, s *goquery.Selection) string {
			return cleanText(s.Text())
		}), ", "),
		WorkingDays:  cleanText(doc.Find(".object_content--right-list.object_content--timetable .object_content-one").Text()),
		WorkingHours: cleanText(doc.Find(".object_content--right-list.object_content--timetable .object_content-too").Text()),
	}

	return detail, nil
}

func main() {
	fileBytes, err := ioutil.ReadFile("sights_data.txt")
	if err != nil {
		log.Fatal("Error reading JSON file: ", err)
	}

	var sights []struct {
		Link string `json:"link"`
	}
	if err := json.Unmarshal(fileBytes, &sights); err != nil {
		log.Fatal("Error unmarshaling JSON: ", err)
	}

	csvFile, err := os.Create("sight_details.csv")
	if err != nil {
		log.Fatal("Error creating CSV file: ", err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write the header
	header := []string{"Name", "Address", "Phone Number", "Content Info", "Bus Numbers", "Working Days", "Working Hours"}
	if err := writer.Write(header); err != nil {
		log.Fatal("Error writing header to CSV: ", err)
	}

	for _, sight := range sights {
		detail, err := fetchSightDetails(strings.TrimSpace(sight.Link))
		if err != nil {
			log.Printf("Error fetching sight details for URL %s: %v\n", sight.Link, err)
			continue
		}

		record := []string{detail.Name, detail.Address, detail.PhoneNumber, detail.ContentInfo, detail.BusNumbers, detail.WorkingDays, detail.WorkingHours}
		if err := writer.Write(record); err != nil {
			log.Fatal("Error writing record to CSV: ", err)
		}
	}

	fmt.Println("Sight details saved to 'sight_details.csv'")
}
