package service

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func buildClient() *http.Client {
	client := &http.Client{
	  Timeout: time.Second * 10,
	}

	return client
}

func buildRequest(method string, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)

	if err != nil {
		request = nil
		return nil, err
	}

	return request, nil
}

func httpGET(url string, userAgent string) (*http.Response, error) {
	client := buildClient()
	request, err := buildRequest("GET", url, nil)

	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Add("Accept-Language", "fr-fr")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", userAgent)

	if err != nil {
		client = nil
		request = nil
		return nil, err
	}

	response, err := client.Do(request)
	request = nil
	client = nil

	if response.StatusCode != 200 {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)
		return nil, errors.New("status was not 200")
	}

	return response, err
}
