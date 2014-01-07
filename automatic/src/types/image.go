package types

type Image struct {
	URL     string `json:"url"`
	Height  int    `json:"pixelHeight"`
	Width   int    `json:"pixelWidth"`
	Caption string `json:"caption"`
	Primary bool   `json:"primary"`
}
