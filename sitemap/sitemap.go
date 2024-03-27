package sitemap

import (
	"encoding/xml"
	"log"
	"net/http"
	"net/url"
	"slices"
	"strings"

	"github.com/Acs176/html-parser-go/parser"
	"golang.org/x/net/html"
)

type SitemapBuilder struct {
	baseUrl      string
	VisitedLinks []string
}

func New(url string) *SitemapBuilder {
	return &SitemapBuilder{url, make([]string, 0)}
}

func (sb *SitemapBuilder) BuildSitemap() []byte {
	// you get a top node / a response body

	// you parse it and get a list of the links

	// you filter only the links that belong to the same domain
	// verify if link has absolute path or relative
	// add the domain to the relative paths

	// you GET them one by one and add them to a list of visited links

	// if a link is already visited, discard it

	resp, err := http.Get(sb.baseUrl)

	if err != nil {
		log.Fatalln(err)
	}

	topNode, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	links := parser.ParseHtml(topNode)
	links = sb.prepareLinks(links)
	sb.crawlLinks(links)
	xmlObject := buildXml(sb.VisitedLinks)

	output, err := xml.MarshalIndent(xmlObject, "", " ")
	if err != nil {
		panic(err)
	}

	return output
}

func (sb *SitemapBuilder) prepareLinks(links []*parser.Link) []*parser.Link {
	var newList = make([]*parser.Link, 0)
	base, err := url.Parse(sb.baseUrl)
	if err != nil {
		log.Fatal(err)
	}

	for _, link := range links {
		if strings.Contains(link.Url, "mailto") {
			continue
		}
		if (strings.Contains(link.Url, "http://") || strings.Contains(link.Url, "https://")) && !strings.Contains(link.Url, sb.baseUrl) {
			continue
		}
		if !strings.Contains(link.Url, "http://") && !strings.Contains(link.Url, "https://") {
			relativeUrl, err := url.Parse(link.Url)
			if err != nil {
				log.Fatal(err)
			}
			absoluteUrl := base.ResolveReference(relativeUrl)
			link.Url = absoluteUrl.String()
		}
		newList = append(newList, link)

	}
	return newList

}

func (sb *SitemapBuilder) crawlLinks(links []*parser.Link) {

	for i := 0; i < len(links); i++ {
		if slices.Contains(sb.VisitedLinks, links[i].Url) {
			continue
		}
		sb.VisitedLinks = append(sb.VisitedLinks, links[i].Url)
		resp, err := http.Get(links[i].Url)
		if err != nil {
			continue
		}
		topNode, err := html.Parse(resp.Body)
		if err != nil {
			continue
		}
		extractedLinks := parser.ParseHtml(topNode)
		extractedLinks = sb.prepareLinks(extractedLinks)
		links = sb.appendUnvisitedLinks(links, extractedLinks)

	}
}

func (sb *SitemapBuilder) appendUnvisitedLinks(ogList []*parser.Link, extracted []*parser.Link) []*parser.Link {
	for _, link := range extracted {
		if !slices.Contains(sb.VisitedLinks, link.Url) {
			ogList = append(ogList, link)
		}
	}
	return ogList
}

func buildXml(links []string) URLSet {
	urlSet := URLSet{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  make([]URL, 0),
	}

	for _, link := range links {
		urlSet.AddUrl(link)
	}
	return urlSet
}
