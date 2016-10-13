package kyriejson

import (
	"fmt"
)

type counterAct int

const (
	counterUrl counterAct = iota
	counterHttpPass
	counterJsonPass
	counterHttpFail
	counterJsonFail
)

type counter struct {
	urls     int
	httpPass int
	jsonPass int
	httpFail int
	jsonFail int
	ch       chan counterAct
}

func newCounter() *counter {
	c := &counter{
		ch: make(chan counterAct),
	}
	go c.run()
	return c
}

func (c *counter) run() {
	for p := range c.ch {
		switch p {
		case counterUrl:
			c.urls++
		case counterHttpPass:
			c.httpPass++
		case counterJsonPass:
			c.jsonPass++
		case counterHttpFail:
			c.httpFail++
		case counterJsonFail:
			c.jsonFail++
		}
	}
}

func (c *counter) total() int {
	return c.httpPass + c.httpFail + c.jsonPass + c.jsonFail
}

func (c *counter) fails() int {
	return c.httpFail + c.jsonFail
}

func (c *counter) print() {
	f := c.fails()
	res := "FAIL"
	if f == 0 {
		res = "PASS"
	}
	fmt.Printf(
		"===========================================================\n"+
			"%s - %d tests run, %d urls; %d test failed\n"+
			"%d/%d HTTP test failed, %d/%d JSON tests failed\n",
		res, c.total(), c.urls, c.httpFail+c.jsonFail,
		c.httpFail, c.httpPass+c.httpFail,
		c.jsonFail, c.jsonPass+c.jsonFail)
}

func (c *counter) startedUrl() {
	c.ch <- counterUrl
}

func (c *counter) passedHttp() {
	c.ch <- counterHttpPass
}

func (c *counter) failedHttp() {
	c.ch <- counterHttpFail
}

func (c *counter) passedJson() {
	c.ch <- counterJsonPass
}

func (c *counter) failedJson() {
	c.ch <- counterJsonFail
}

func (c *counter) close() {
	close(c.ch)
}
