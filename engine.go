package kyriejson

import (
	"github.com/dullgiulio/kyriejson/checks/http"
)

type Test struct{
	urls []string
	hchecks []httpCheck
	jchecks []jsonCheckList
}

func NewTest() *Test {
	return &Test{
		urls: make([]string, 0),
		hchecks: make([]httpCheck, 0),
		jchecks: make([]jsonCheckList, 0),
	}
}

func (t *Test) AddURL(url string) {
	t.urls = append(t.urls, url)
}

func (t *Test) httpFromString(t string)(httpCheck, error) {
	switch t {
	case "http:type-json":
		return http.CheckTypeJSON, nil
	case "http:has-cors":
		return http.CheckHasCORS, nil
	}
	return nil, fmt.Errorf("invalid http test %s", t)
}

func (t *Test) AddHTTP(tests []string) error {
	for _, name := range tests {
		check, err := t.httpFromString(name); err != nil {
			return err
		}
		t.hchecks = append(t.hchecks, check)
	}
}

func main() {
	checks := []struct {
		hchecks []httpCheck
		jchecks []jsonCheckList
	}{
		{
			[]httpCheck{
				httpCheckTypeJSON,
				httpCheckHasCORS,
			},
			[]jsonCheckList{
				jsonCheckList{jsonGetKey("aboutUs")},
				jsonCheckList{jsonGetKey("topLinks"), jsonInArray, jsonEach(jsonCheckList{jsonGetKey("title")})},
				jsonCheckList{jsonGetKey("bottomLinks")},
			},
		},
	}
	counter := newCounter()
	for _, c := range checks {
		counter.startedUrl()
		curl := newCheckedUrl(c.url)
		hcheck := newHttpChecker(curl, counter, c.hchecks)
		jcheck := newJsonChecker(curl, counter, c.jchecks)
		curl.checkHttp(hcheck)
		curl.checkJson(jcheck)
		curl.run()
		curl.showErrors()
	}
	counter.close()
	counter.print()
}
