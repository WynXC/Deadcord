package core

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

func HarvestFromURL(url string, timeout int) ([]string, error) {
	proxy_match := regexp.MustCompile(`(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?).){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?):([0-9]){1,4}`)
	proxy_client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	response, err := proxy_client.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	proxies := proxy_match.FindAllString(string(body), -1)
	return proxies, nil
}
