package types

type Video struct {
	URL     string `json:"url"`
	Height  int    `json:"pixelHeight"`
	Width   int    `json:"pixelWidth"`
	Primary bool   `json:"primary"`
}
