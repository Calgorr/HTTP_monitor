package monitor

import (
	"net/http"
	"net/url"
)

type Resp struct {
	StatusCode int
	err        string
}

func SendRequest(urlString string, client http.Client, result chan<- *Resp) {
	var response *Resp
	req := http.Request{}
	req.URL, _ = url.Parse(urlString)
	req.Method = "GET"
	resp, err := client.Do(&req)
	if err != nil {
		response.StatusCode = 0
	}
	response.StatusCode = resp.StatusCode
	result <- response
}
