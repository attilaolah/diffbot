package diffbot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"./types"
)

const (
	apiArticle = apiRoot + "/article"
)

func (c *Client) Article(url_ string) (a *types.Article, err error) {
	u, _ := url.Parse(apiArticle)
	q := u.Query()
	q.Set("token", c.APIKey)
	q.Set("url", url_)
	q.Set("fields", "*") // debug
	u.RawQuery = q.Encode()

	r, err := c.Client.Get(u.String())
	if err != nil {
		return
	}
	if r.StatusCode != http.StatusOK {
		err = fmt.Errorf("diffbot: api returned %s", r.Status)
		return
	}
	defer r.Body.Close()
	a = new(types.Article)
	err = json.NewDecoder(r.Body).Decode(a)
	return
}
