package diffbot

import (
	"net/http"
)

type Client struct {
	Client http.Client
}

var client Client