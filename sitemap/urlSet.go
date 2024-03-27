package sitemap

import "encoding/xml"

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

type URL struct {
	Loc string `xml:"loc"`
}

func (us *URLSet) AddUrl(url string) {
	urlObject := URL{url}
	us.URLs = append(us.URLs, urlObject)
}
