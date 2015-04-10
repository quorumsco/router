package iogo

import (
	"io"
	"net/http"
)

type IogoContext struct {
	C      map[interface{}]interface{}
	Params []param
	req    *http.Request
}

type iogoContextReadCloser struct {
	io.ReadCloser
	context *IogoContext
}

func setNewContext(req *http.Request) *IogoContext {
	c, ok := req.Body.(iogoContextReadCloser)
	if !ok {
		c = iogoContextReadCloser{
			ReadCloser: req.Body,
			context:    &IogoContext{C: make(map[interface{}]interface{})},
		}
		req.Body = c
	}
	return c.context
}

func GetContext(req *http.Request) *IogoContext {
	return req.Body.(iogoContextReadCloser).context
}

func (c *IogoContext) GetParam(name string) string {
	for _, e := range c.Params {
		if e.key == name {
			return e.value
		}
	}
	return ""
}

func (c *IogoContext) SetParams(params []param) {
	c.Params = params
}
