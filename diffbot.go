// Package diffbot provides a wrapper for the Diffbot API.
package diffbot

import (
	"./types"
)

const apiRoot = "http://api.diffbot.com/v2"

func Article(key, url string) (a *types.Article, err error) {
	c := Client{APIKey: key}
	return c.Article(url)
}
