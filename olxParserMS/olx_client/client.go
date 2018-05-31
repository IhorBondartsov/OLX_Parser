package olx_client

import (
	"fmt"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/dateParser"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/entities"
	"github.com/PuerkitoBio/goquery"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strings"
)

type OLXClient struct {
	httpClient *http.Client
}

type Result struct {
	URL  string
	Date string
}

func GetDocumentByUrl(url string) *goquery.Document {
	log.Info("[GetDocumentByUrl]", url)

	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return nil
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return doc
}

func GetAdvertisements(doc *goquery.Document) []entities.Advertisement{
	var advrtmnts []entities.Advertisement
	if doc == nil {
		return advrtmnts
	}
	// Load the HTML document
	doc.Find("#offers_table").Each(func(i int, s *goquery.Selection) {
		s.Find(".wrap").Each(func(i int, s *goquery.Selection) {

			advrtmnt := entities.Advertisement{}
			title := s.Find("h3")
			url, _ := title.Find("a").Attr("href")
			time := s.Find(".space").Eq(2).Find("p").Last().Text()

			advrtmnt.Title = strings.TrimSpace(title.Text())
			advrtmnt.URL = strings.TrimSpace(url)
			advrtmnt.Time = dateParser.ParseTime(time)

			advrtmnts = append(advrtmnts, advrtmnt)
		})
	})
	return advrtmnts
}

func GetHTMLPages(url string) []entities.Advertisement {
	log.Info("[GetHTMLPages]", url)

	var advrtmnts []entities.Advertisement

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

	advrtmnts = append(advrtmnts, GetAdvertisements(doc)...)

	doc.Find(".pager").Each(func(i int, s *goquery.Selection) {
		s.Find(".next").Each(func(i int, s *goquery.Selection) {
			url, _ := s.Find("a").Attr("href")
			advrtmnts = append(advrtmnts, GetAdvertisements(GetDocumentByUrl(url))...)
		})

	})
	fmt.Println(len(advrtmnts))
	return advrtmnts
}

