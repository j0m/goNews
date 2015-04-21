// http://www.w3schools.com/rss/default.asp
// http://www.tutorialspoint.com/rss/what-is-atom.htm
// http://stackoverflow.com/questions/16309944/atom-to-rss-feeds-converter
package feedParser

import (
	"encoding/xml"
	"fmt"
	"html/template"
)

// http://www.w3schools.com/rss/rss_syntax.asp
// http://www.w3schools.com/rss/rss_channel.asp
type Rss2 struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	// Required
	Title       string `xml:"channel>title"`
	Link        string `xml:"channel>link"`
	Description string `xml:"channel>description"`
	// Optional
	PubDate  string     `xml:"channel>pubDate"`
	ItemList []Rss2Item `xml:"channel>item"`
}

// http://www.w3schools.com/rss/rss_item.asp
// http://stackoverflow.com/questions/7220670/difference-between-description-and-contentencoded-tags-in-rss2
// https://groups.google.com/d/topic/golang-nuts/uBMo1BpaQCM
type Rss2Item struct {
	// Required
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Description template.HTML `xml:"description"`
	// Optional
	Content  template.HTML `xml:"encoded"`
	PubDate  string        `xml:"pubDate"`
	Comments string        `xml:"comments"`
}

func (r *Rss2) Parse(content []byte) {
	err := xml.Unmarshal(content, r)

	if err != nil {
		panic(err)
	}

	for _, item := range r.ItemList {
		fmt.Println(item)
		fmt.Println("----------------------------------------------------")
	}
}

func NewRssFeed() *Rss2 {
	return &Rss2{}
}
