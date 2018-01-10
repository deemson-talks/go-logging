package main

import (
	"net/http"
	"io/ioutil"
	"github.com/pkg/errors"
	"log"
	"fmt"
)

func Download(url string) ([]byte, error) {
	errors.New()
	r, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "get url: %s", url)
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "writing file")
	}
	return body, nil
}

func main() {
	data, err := Download("https://non-existent-url")
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	fmt.Println(data)
}
