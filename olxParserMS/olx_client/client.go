package olx_client

import (
	"net/http"
	log "github.com/Sirupsen/logrus"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/entities"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/dateParser"
	"strings"
)
type OLXClient struct{
	httpClient *http.Client
}

type Result struct{
	URL string
	Date string
}

func GetHTMLPage(url string)[]entities.Advertisement	{
	log.Info("[GetHTMLPage]", url)

	var advrtmnts  []entities.Advertisement

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
		s.Find("tbody").Each(func(i int, s *goquery.Selection){

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
	fmt.Println(advrtmnts)
	return advrtmnts
	}

	//fmt.Println(doc.Find("#offers_table").Nodes[0].Namespace)
