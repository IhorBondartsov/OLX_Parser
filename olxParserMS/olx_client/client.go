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
	fmt.Println(doc.Find("tbody").)
}