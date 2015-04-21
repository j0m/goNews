// http://www.tutorialspoint.com/rss/what-is-atom.htm
// http://stackoverflow.com/questions/16309944/atom-to-rss-feeds-converter
package feedParser

import (
	"encoding/xml"
	"fmt"
)

// http://en.wikipedia.org/wiki/Atom_%28standard%29
// http://golang.org/src/pkg/encoding/xml/
type Atom1 struct {
	XMLName   xml.Name     `xml:"http://www.w3.org/2005/Atom feed"`
	Title     string       `xml:"title"`
	Subtitle  string       `xml:"subtitle"`
	Id        string       `xml:"id"`
	Updated   string       `xml:"updated"`
	Rights    string       `xml:"rights"`
	Link      Atom1Link    `xml:"link"`
	Author    Atom1Author  `xml:"author"`
	EntryList []Atom1Entry `xml:"entry"`
}

type Atom1Link struct {
	Href string `xml:"href,attr"`
}

type Atom1Author struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
}

type Atom1Entry struct {
	Title   string      `xml:"title"`
	Summary string      `xml:"summary"`
	Content string      `xml:"content"`
	Id      string      `xml:"id"`
	Updated string      `xml:"updated"`
	Link    Atom1Link   `xml:"link"`
	Author  Atom1Author `xml:"author"`
}

func (a *Atom1) Parse(content []byte) {
	err := xml.Unmarshal(content, a)

	if err != nil {
		panic(err)
	}

	for _, entry := range a.EntryList {
		fmt.Println(entry)
	}
}

func NewAtomFeed() *Atom1 {
	return &Atom1{}
}
