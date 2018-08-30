package blkchain

import (
	"io"
	"io/ioutil"
	"net/http"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: 10,
		},
	}
}

func RPCRequest(method, requestUrl string, requestBody io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, requestUrl, requestBody)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
