package olx_client

import (
	"net/http"
	log "github.com/Sirupsen/logrus"
	"github.com/PuerkitoBio/goquery"
	"fmt"
)
type OLXClient struct{
	httpClient *http.Client
}

type Result struct{
	URL string
	Date string
}

func GetHTMLPage(url string) 	{
	log.Info("[GetHTMLPage]", url)

	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("#offers_table").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		s.Find("tbody").Each(func(i int, s *goquery.Selection){
			title := s.Find("h3")
			url, ok := title.Find("a").Attr("href")
			fmt.Printf("Review %d: %s - %s  - %s \n", i, title.Text(), url, ok)

			time := s.Find(".space").Eq(2).Find("p").Last().Text()
			fmt.Println("--------------> ", time)
			})

		})

	}

	//fmt.Println(doc.Find("#offers_table").Nodes[0].Namespace)
