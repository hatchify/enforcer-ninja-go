package sdk

// UsageEntry represents a date entry for the get usage response
type UsageEntry struct {
	Date  string `json:"date"`
	Usage uint64 `json:"usage"`
}
