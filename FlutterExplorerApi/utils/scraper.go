package utils

import (
	"flutterexplorer/models"
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func PackagesScraper(name string) ([]models.Package, error) {
	var packages []models.Package
	var platform []string
	var err error

	c := colly.NewCollector(colly.Async(true))
	c.OnHTML("body > main > div > div.search-results > div.packages > div.packages-item", func(element *colly.HTMLElement) {

		element.ForEach("div.-pub-tag-badge > a.tag-badge-sub", func(i int, element *colly.HTMLElement) {
			platform = append(platform, element.Text)
		})

		packages = append(packages, models.Package{
			Name:        element.ChildText("div.packages-header > h3"),
			Description: element.ChildText("p.packages-description"),
			Link:        element.ChildAttr("div.packages-header > h3 > a", "href"),
			Scores: models.Scores{
				Likes:      element.ChildText("a.packages-scores > div.packages-score-like > div.packages-score-value"),
				PubPoints:  element.ChildText("a.packages-scores > div.packages-score-health > div.packages-score-value"),
				Popularity: element.ChildText("a.packages-scores > div.packages-score-popularity > div.packages-score-value"),
			},
			Metadata: models.Metadata{
				Version: element.ChildText("p.packages-metadata > span:nth-child(1)"),
			},
			Badge: models.Badge{
				Platform: platform,
			},
		})
	})

	c.OnError(func(response *colly.Response, e error) {
		err = e
	})

	c.Visit(fmt.Sprintf("https://pub.dev/packages?q=%s", name))
	c.Wait()
	return packages, err
}

func WidgetScraper() ([]models.Widget, error) {
	var widgets []models.Widget
	var description []string
	var name []string
	var link []string
	var err error

	c := colly.NewCollector(colly.Async(true))
	c.OnHTML("dl > dd", func(e *colly.HTMLElement) {
		description = append(description, strings.ReplaceAll(strings.TrimSpace(e.Text), "\n", ""))
	})

	c.OnHTML("dt", func(e *colly.HTMLElement) {
		name = append(name, strings.ReplaceAll(strings.TrimSpace(e.Text), "\n", ""))
		link = append(link, e.ChildAttr("span > a", "href"))
	})

	c.OnScraped(func(response *colly.Response) {
		for i := 0; i < len(name); i++ {
			widgets = append(widgets, models.Widget{Name: name[i], Description: description[i], Link: link[i]})
		}
	})

	c.OnError(func(response *colly.Response, e error) {
		err = e
	})

	c.Visit("https://api.flutter.dev/flutter/widgets/widgets-library.html")
	c.Wait()

	return widgets, err
}
