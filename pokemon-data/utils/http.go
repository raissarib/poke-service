package utils

import (
	"io"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 5,
}

func Get(url string) ([]byte, error) {
	res, err := client.Get(url)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	return io.ReadAll(res.Body)
}
