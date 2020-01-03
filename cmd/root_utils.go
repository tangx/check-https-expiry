package cmd

import (
	"net/http"
	"time"
)

func httpGet(url string) (resp *http.Response, err error) {
	cli := http.Client{
		Timeout: 3 * time.Second,
	}
	return cli.Get(url)
}
