package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

type Quote struct {
	Text   string   `json:"text"`
	Author string   `json:"author"`
	Tags   []string `json:"tags"`
}

func main() {
	
	http.HandleFunc("/quotes", func(w http.ResponseWriter, r *http.Request) {
		citations := scrapeQuotes(100) // Retrieve exactly 100 quotes

		// Return the quotes in JSON format
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(citations)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Function for scraping quotes
func scrapeQuotes(limit int) []Quote {
	var quotes []Quote
	page := 1
	url := "https://quotes.toscrape.com/page/%d/"

	// Scrape pages until we have 100 quotes 
	for {
		// Build the URL for the current pages
		currentURL := fmt.Sprintf(url, page)
		fmt.Printf("Scraping page %d: %s\n", page, currentURL)

		// fetch the page content
		res, err := http.Get(currentURL)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		// Ensure the page was retrieved successfully
		if res.StatusCode != 200 {
			break // If the page doesn't exist, exit the loop
		}

		// Parse the page with goquery
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Extract quotes for each page
		doc.Find(".quote").Each(func(i int, s *goquery.Selection) {
			quoteText := s.Find(".text").Text()
			author := s.Find(".author").Text()
			var tags []string

			// Extract tags for each quote
			s.Find(".tags .tag").Each(func(i int, tag *goquery.Selection) {
				tags = append(tags, tag.Text())
			})

			// Add the quote, author, and tags to the list 
			quotes = append(quotes, Quote{Text: quoteText, Author: author, Tags: tags})
		})

		// check if we have reached the limit of 100 quotes 
		if len(quotes) >= limit {
			break
		}

		// Move to the next page
		page++

		// If the last page is reached (check for the absence of the "Next" button), stop
		if doc.Find(".next").Length() == 0 {
			break
		}
	}

	// Limit the quotes to 100 if more than 100 were retrieved
	if len(quotes) > limit {
		quotes = quotes[:limit]
	}

	
	
	return quotes
}
