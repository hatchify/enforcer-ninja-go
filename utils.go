package sdk

import (
	"net/url"
	"path"
)

func getURL(endpoint string, query url.Values) (u *url.URL, err error) {
	if u, err = url.Parse(Host); err != nil {
		return
	}

	u.Path = path.Join(APIPath, endpoint)

	if query != nil {
		u.RawQuery = query.Encode()
	}

	return
}

func getViewQuery(v View) (q url.Values) {
	q = url.Values{}
	q.Set("view", string(v))
	return
}

// Config represents a basic configuration
type Config struct {
	APIKey string `toml:"apiKey" json:"apiKey"`
}
