package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errlist interface {
	add(error)
}

type checkedUrl struct {
	errs     []error
	url      string
	response interface{}
	hcheck   *httpChecker
}

func newCheckedUrl(url string) *checkedUrl {
	return &checkedUrl{
		url:      url,
		errs:     make([]error, 0),
		response: make(map[string]interface{}),
	}
}

func (c *checkedUrl) checkHttp(hcheck *httpChecker) {
	c.hcheck = hcheck
}

func (c *checkedUrl) add(e error) {
	c.errs = append(c.errs, e)
}

func (c *checkedUrl) load() error {
	r, err := http.Get(c.url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	c.hcheck.run(r)

	return json.NewDecoder(r.Body).Decode(&c.response)
}

func (c *checkedUrl) showErrors() {
	for i := range c.errs {
		fmt.Printf("E: %s: %s\n", c.url, c.errs[i])
	}
}

func (c *checkedUrl) run() {
	if err := c.load(); err != nil {
		c.add(err)
		return
	}
}
