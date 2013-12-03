package main

import (
	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
	"time"
)

// Only enqueue the root and paths beginning with an "a"
var rxOk = regexp.MustCompile(`http://duckduckgo\.com(/a.*)?$`)

// Create the Extender implementation, based on the gocrawl-provided DefaultExtender,
// because we don't want/need to override all methods.
type ExampleExtender struct {
	gocrawl.DefaultExtender // Will use the default implementation of all but Visit() and Filter()
}

// Override Visit for our need.
func (this *ExampleExtender) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	// Use the goquery document or res.Body to manipulate the data
	// ...

	// Return nil and true - let gocrawl find the links

	return nil, true
}

// Override Filter for our need.
func (this *ExampleExtender) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	return !isVisited && rxOk.MatchString(ctx.NormalizedURL().String())
}

func ExampleCrawl() {
	// Set custom options
	opts := gocrawl.NewOptions(new(ExampleExtender))
	opts.CrawlDelay = 1 * time.Second
	opts.LogFlags = gocrawl.LogAll

	// Play nice with ddgo when running the test!
	opts.MaxVisits = 2

	// Create crawler and start at root of duckduckgo
	c := gocrawl.NewCrawlerWithOptions(opts)
	c.Run("https://duckduckgo.com/")
}

func main() {
	ExampleCrawl()
}
