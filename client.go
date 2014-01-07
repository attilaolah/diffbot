package diffbot

import (
	"net/http"
)

type Client struct {
	APIKey string
	Client http.Client
}
