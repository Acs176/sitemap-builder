# Sitemap Builder

This project provides a tool for generating sitemaps for websites.

## Table of Contents
- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)

## Overview
The Sitemap Builder is a Go package designed to create sitemaps for websites. It crawls through the provided URL, extracts links, filters them based on the domain, and generates an XML sitemap containing all the visited links.

## Installation
To use this package in your Go project, you can install it via:

```bash
go get github.com/Acs176/sitemap
```

## Usage
Below is an example of how to use the Sitemap Builder:

```go
package main

import (
    "fmt"
    "github.com/Acs176/sitemap"
)

func main() {
    builder := sitemap.New("https://example.com")
    xmlSitemap := builder.BuildSitemap()
    fmt.Println(string(xmlSitemap))
}
```
