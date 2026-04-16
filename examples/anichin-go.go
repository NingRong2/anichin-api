// Anichin API (Dracin API) - Go Example
//
// Jalankan: go run examples/anichin-go.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	AnichinBase  = "https://api.anichin.bio"
	AnichinToken = "TRIAL-ANICHIN-2026"
)

type AnichinClient struct{ http *http.Client }

func NewClient() *AnichinClient {
	return &AnichinClient{http: &http.Client{Timeout: 30 * time.Second}}
}

func (c *AnichinClient) Get(source, path string, params map[string]string) (map[string]any, error) {
	u, _ := url.Parse(fmt.Sprintf("%s/%s/%s", AnichinBase, source, path))
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("X-API-Key", AnichinToken)
	req.Header.Set("User-Agent", "Mozilla/5.0 (anichin-go-example)")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	c := NewClient()

	trending, err := c.Get("dramabox", "trending", nil)
	if err != nil {
		panic(err)
	}
	items := trending["items"].([]any)
	fmt.Printf("Anichin trending (DramaBox): %d drama\n", len(items))

	first := items[0].(map[string]any)
	fmt.Printf("#1 -> %s\n", first["title"])

	ep, _ := c.Get("dramabox", "episode", map[string]string{
		"id": fmt.Sprint(first["id"]),
		"ep": "1",
	})
	fmt.Printf("Episode 1 video: %s\n", ep["videoUrl"])
}
