package api

import (
	"io"
	"net/http"
	"strings"
)

type Upyun struct {
	authorization string
	host          string
}

type Pager struct {
	Since any `json:"since"`
	Max   int `json:"max"`
	Limit int `json:"limit"`
}

func (u Upyun) Get(url string, body io.Reader) ([]byte, error) {
	return u.Request("GET", url, body)
}

func (u Upyun) Request(method, url string, body io.Reader) ([]byte, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", u.authorization)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return io.ReadAll(response.Body)
}

type UpyunConfig struct {
	Authorization string
	Host          string
}

func NewUpyun(config UpyunConfig) Upyun {
	if !strings.HasPrefix(config.Authorization, "Bearer") {
		config.Authorization = "Bearer " + config.Authorization
	}
	if len(config.Host) == 0 {
		config.Host = "https://api.upyun.com"
	}
	return Upyun{
		authorization: config.Authorization,
		host:          config.Host,
	}
}
