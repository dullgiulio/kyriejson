package kyriejson

import (
	"net/http"
)

type httpCheck func(*http.Response) error

type httpChecker struct {
	errors  errlist
	counter *counter
	checks  []httpCheck
}

func newHttpChecker(errors errlist, counter *counter, checks []httpCheck) *httpChecker {
	return &httpChecker{
		errors:  errors,
		counter: counter,
		checks:  checks,
	}
}

func (h *httpChecker) run(resp *http.Response) {
	for _, check := range h.checks {
		if err := check(resp); err != nil {
			h.errors.add(err)
			h.counter.failedHttp()
			continue
		}
		h.counter.passedHttp()
	}
}

type jsonCheck func(jsonData) (jsonData, error)

type jsonCheckList []jsonCheck

func (js jsonCheckList) run(data jsonData) error {
	for _, check := range js {
		d, err := check(data)
		if err != nil {
			return err
		}
		if d == nil {
			return nil
		}
		data = d
	}
	return nil
}

type jsonChecker struct {
	errors  errlist
	counter *counter
	checks  []jsonCheckList
}

func newJsonChecker(errors errlist, counter *counter, checks []jsonCheckList) *jsonChecker {
	return &jsonChecker{
		errors:  errors,
		counter: counter,
		checks:  checks,
	}
}

func (j *jsonChecker) run(data jsonData) {
	for _, list := range j.checks {
		if err := list.run(data); err != nil {
			j.errors.add(err)
			j.counter.failedJson()
			continue
		}
		j.counter.passedJson()
	}
}
