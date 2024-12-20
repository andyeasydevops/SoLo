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
		citations := scrapeQuotes(100) // Récupérer exactement 100 citations

		// Retourner les citations en format JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(citations)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Fonction de scraping des citations
func scrapeQuotes(limit int) []Quote {
	var quotes []Quote
	page := 1
	url := "https://quotes.toscrape.com/page/%d/"

	// Scraper les pages jusqu'à ce que nous ayons 100 citations
	for {
		// Construire l'URL pour la page actuelle
		currentURL := fmt.Sprintf(url, page)
		fmt.Printf("Scraping page %d: %s\n", page, currentURL)

		// Récupérer le contenu de la page
		res, err := http.Get(currentURL)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		// Vérifier que la page a été récupérée avec succès
		if res.StatusCode != 200 {
			break // Si la page n'existe pas, sortir de la boucle
		}

		// Parser la page avec goquery
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Récupérer les citations de la page
		doc.Find(".quote").Each(func(i int, s *goquery.Selection) {
			quoteText := s.Find(".text").Text()
			author := s.Find(".author").Text()
			var tags []string

			// Récupérer les tags pour chaque citation
			s.Find(".tags .tag").Each(func(i int, tag *goquery.Selection) {
				tags = append(tags, tag.Text())
			})

			// Ajouter la citation, l'auteur et les tags à la liste
			quotes = append(quotes, Quote{Text: quoteText, Author: author, Tags: tags})
		})

		// Vérifier si on a atteint la limite de 100 citations
		if len(quotes) >= limit {
			break
		}

		// Passer à la page suivante
		page++

		// Si on atteint la dernière page (vérifier l'absence du bouton "Next"), arrêter
		if doc.Find(".next").Length() == 0 {
			break
		}
	}

	// Limiter les citations à 100 si plus de 100 ont été récupérées
	if len(quotes) > limit {
		quotes = quotes[:limit]
	}

	
	
	return quotes
}