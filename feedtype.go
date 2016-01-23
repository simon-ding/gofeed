package feed

import (
	"encoding/xml"
	"strings"
)

type FeedType int

const (
	FeedTypeUnknown FeedType = iota
	FeedTypeAtom
	FeedTypeRSS
)

// Performs a simple check on the feed to determine
// the feed type.
func DetectFeedType(feed string) FeedType {
	decoder := xml.NewDecoder(strings.NewReader(feed))
	decoder.Strict = false
	for {
		token, err := decoder.Token()

		if err != nil || token == nil {
			return FeedTypeUnknown
		}

		switch t := token.(type) {
		case xml.StartElement:
			name := strings.ToLower(t.Name.Local)
			switch name {
			case "rdf":
				return FeedTypeRSS
			case "rss":
				return FeedTypeRSS
			case "feed":
				return FeedTypeAtom
			}
		}
	}
}