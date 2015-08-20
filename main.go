package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/daviddengcn/go-colortext"
	"github.com/mattn/go-runewidth"
)

func main() {
	lang := ""
	switch len(os.Args) {
	case 1:
	case 2:
		lang = os.Args[1]
	default:
		fmt.Fprintf(os.Stderr, "usage of %s: [language]", os.Args[0])
		os.Exit(1)
	}
	uri := "https://github.com/trending"
	if lang != "" {
		uri += "?l=" + lang
	}
	doc, err := goquery.NewDocument(uri)
	if err != nil {
		fmt.Printf("%s: %s\n", os.Args[0], err)
		os.Exit(1)
	}

	doc.Find(".repo-list-item").Each(func(_ int, s *goquery.Selection) {
		if href, ok := s.Find(".repo-list-name a").First().Attr("href"); ok {
			if strings.HasPrefix(href, "/") {
				href = "https://github.com" + href
			}
			desc := strings.TrimSpace(s.Find(".repo-list-description").First().Text())
			desc = runewidth.Wrap(desc, 76)
			desc = "  " + strings.Replace(desc, "\n", "\n  ", -1)

			ct.ChangeColor(ct.Yellow, false, ct.None, false)
			fmt.Println(href)
			ct.ChangeColor(ct.Green, false, ct.None, false)
			fmt.Println(desc)
			ct.ResetColor()
		}
	})
}
