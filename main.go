package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/Acs176/sitemap-builder/sitemap"
)

func main() {

	var sb = sitemap.New("https://www.calhoun.io/")
	output := sb.BuildSitemap()
	file, err := os.Create("sitemap.xml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write([]byte(xml.Header))
	file.Write(output)
	fmt.Println("Completed generating sitemap!!")
}
