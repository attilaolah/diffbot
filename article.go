package diffbot

import (
	"encoding/json"

	"./types"
)

const (
	apiArticle = apiRoot + "/article"
)

func (c *Client) Article(url string) (a *types.Article, err error) {
	r, err := c.Client.Get(apiArticle)
	if err != nil {
		return
	}
	defer r.Body.Close()
	a = new(types.Article)
	err = json.NewDecoder(r.Body).Decode(a)
	return
}
