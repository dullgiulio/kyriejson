package kyriejson

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type jsonData interface{}

type errlist interface {
	add(error)
}

type checkedUrl struct {
	errs   []error
	url    string
	data   jsonData
	hcheck *httpChecker
	jcheck *jsonChecker
}

func newCheckedUrl(url string) *checkedUrl {
	return &checkedUrl{
		url:  url,
		errs: make([]error, 0),
		data: jsonData(make(map[string]interface{})),
	}
}

func (c *checkedUrl) checkHttp(hcheck *httpChecker) {
	c.hcheck = hcheck
}

func (c *checkedUrl) checkJson(jcheck *jsonChecker) {
	c.jcheck = jcheck
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

	if c.hcheck != nil {
		c.hcheck.run(r)
	}
	if err := json.NewDecoder(r.Body).Decode(&c.data); err != nil {
		return err
	}
	if c.jcheck != nil {
		c.jcheck.run(c.data)
	}
	return nil
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
