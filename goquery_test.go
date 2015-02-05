package main

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"testing"
)

func TestParseHtml(t *testing.T) {
	html := `
<section>
<h1 class="heading">うんこっこ</h1>
<p class="lead">なんかにおうってばよ</p>
</section>
`
	r := strings.NewReader(html)
	doc, _ := goquery.NewDocumentFromReader(r)

	docHeading := doc.Find(".heading").Text()
	heading := "うんこっこ"

	if docHeading != heading {
		t.Errorf("got %v want %v", docHeading, heading)
	}

	docLead := doc.Find(".heading").Text()
	lead := "うんこっこ"

	if docLead != lead {
		t.Errorf("got %v want %v", docLead, lead)
	}
}
