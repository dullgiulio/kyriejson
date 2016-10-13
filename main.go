package main

func main() {
	checks := []struct {
		url     string
		hchecks []httpCheck
		jchecks []jsonCheck
	}{
		{
			"http://dev.kuehne-nagel.int.kn/webservice/links",
			[]httpCheck{
				httpCheckTypeJSON,
				httpCheckHasCORS,
			},
			[]jsonCheck{},
		},
	}
	counter := newCounter()
	for _, c := range checks {
		curl := newCheckedUrl(c.url)
		hcheck := newHttpChecker(curl, counter, c.hchecks)
		curl.checkHttp(hcheck)
		// TODO: jsonChecks
		curl.run()
		curl.showErrors()
	}
	counter.close()
	counter.print()
}
