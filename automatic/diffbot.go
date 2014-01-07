// Package diffbot provides a wrapper for the Diffbot API.
package diffbot

import (
	"./types"
)

const apiRoot = "http://api.diffbot.com/v2"

func Article(url string) (a *types.Article, err error) {
	return client.Article(url)
}
