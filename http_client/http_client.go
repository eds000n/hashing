package http_client

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
)

func DoRequest(url string) ([16]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return [16]byte{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return [16]byte{}, err
	}

	return md5.Sum(body), nil
}
