package link

import "io"

// Represents a link in an html document
type Link struct {
	Href string
	Text string
}

// Take an HTML document and returns a slice of links
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
