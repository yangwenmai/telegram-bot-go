package command

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ListGithubTrending list github trending
func ListGithubTrending(language, since string) string {
	var doc *goquery.Document
	var e error
	result := "\n```\n" + "since:" + since + "\n\n#### " + language + "\n"

	if doc, e = goquery.NewDocument("https://github.com/trending?l=" + language + "&since=" + since); e != nil {
		println("Error:", e.Error())
		panic(language)
	}

	doc.Find("ol.repo-list li").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h3 a").Text()
		description := s.Find("p.col-9").Text()
		url, _ := s.Find("h3 a").Attr("href")
		url = "https://github.com" + url
		var stars = "0"
		var forks = "0"
		s.Find("a.muted-link.mr-3").Each(func(i int, contentSelection *goquery.Selection) {
			if temp, ok := contentSelection.Find("svg").Attr("aria-label"); ok {
				switch temp {
				case "star":
					stars = contentSelection.Text()
				case "fork":
					forks = contentSelection.Text()
				}
			}
		})
		result = result + "* [" + strings.Replace(strings.TrimSpace(title), " ", "", -1) + " (" + strings.TrimSpace(stars) + "star/" + strings.TrimSpace(forks) + "fork)](" + url + ") : " + strings.TrimSpace(description) + "\n"
	})
	result += "```"
	return result
}
