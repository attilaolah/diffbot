package diffbot

import (
	"./types"
)

const (
	apiArticle = apiRoot + "/article"
)

func (c *Client) Article() (a types.Article, err error) {
	// TODO: fetch and decode the article!
	return
}
