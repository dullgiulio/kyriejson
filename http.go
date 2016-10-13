package main

import (
	"errors"
	"net/http"
)

func httpCheckHasCORS(r *http.Response) error {
	if r.Header.Get("Access-Control-Allow-Origin") != "*" {
		return errors.New("CORS header invalid, expected set to star")
	}
	return nil
}

func httpCheckTypeJSON(r *http.Response) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return errors.New("Content-Type header is not valid JSON, application/json")
	}
	return nil
}

// TODO: Implement checks for:
//		 - Cache headers
//		 - ETag is set
