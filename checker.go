package main

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

type jsonCheck func(map[string]interface{}) error

type jsonChecker struct {
	errors  errlist
	counter *counter
	checks  []jsonCheck
}

func newJsonChecker(errors errlist, counter *counter, checks []jsonCheck) *jsonChecker {
	return &jsonChecker{
		errors:  errors,
		counter: counter,
		checks:  checks,
	}
}

func (j *jsonChecker) run(data map[string]interface{}) {
	for _, check := range j.checks {
		if err := check(data); err != nil {
			j.errors.add(err)
			j.counter.failedJson()
			continue
		}
		j.counter.passedJson()
	}
}
