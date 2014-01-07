package types

import (
	"encoding/json"
)

type Article struct {
	URL         string          `json:"url"`
	ResolvedURL string          `json:"resolved_url"`
	Icon        string          `json:"icon"`
	Meta        json.RawMessage `json:"meta" diffbot:"optional"`
	QueryString string          `json:"querystring" diffbot:"optional"`
	Links       []string        `json:"links" diffbot:"optional"`
	Type        string          `json:"type"`
	Title       string          `json:"title"`
	Text        string          `json:"text"`
	HTML        string          `json:"html"`
	Pages       int             `json:"numPages"`
	Date        Time            `json:"date"`
	Author      string          `json:"author"`
	Tags        []string        `json:"tags" diffbot:"optional"`
	Language    string          `json:"humanLanguage" diffbot:"optional"`
	Images      []Image         `json:"images"`
	Videos      []Video         `json:"videos"`

	Stats     Stats      `json:"stats" diffbot:"experimental"`
	Summary   string     `json:"summary" diffbot:"experimental"`
	SuperTags []SuperTag `json:"supertags" diffbot:"experimental"`
}
