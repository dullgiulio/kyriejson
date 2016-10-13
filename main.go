package main

func main() {
	checks := []struct {
		url     string
		hchecks []httpCheck
		jchecks []jsonCheckList
	}{
		{
			"http://dev.kuehne-nagel.int.kn/webservice/links",
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
