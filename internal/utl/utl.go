package utl

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func WanIPv4() (string, error) {
	return wanIP("https://v4.ident.me/")
}

func WanIPv6() (string, error) {
	return wanIP("https://v6.ident.me/")
}

func wanIP(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(body)), nil
}
