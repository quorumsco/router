package router

import (
	"io"
	"net/http"
)

// Context is...
type Context struct {
	Env    map[interface{}]interface{}
	Params []param
	req    *http.Request
}

type iogoContextReadCloser struct {
	io.ReadCloser
	context *Context
}

func setContext(req *http.Request) *Context {
	c, ok := req.Body.(iogoContextReadCloser)
	if !ok {
		c = iogoContextReadCloser{
			ReadCloser: req.Body,
			context:    &Context{Env: make(map[interface{}]interface{})},
		}
		req.Body = c
	}
	return c.context
}

// GetContext is...
func GetContext(req *http.Request) *Context {
	return req.Body.(iogoContextReadCloser).context
}

// GetParam is...
func (c *Context) GetParam(name string) string {
	for _, e := range c.Params {
		if e.name == name {
			return e.value
		}
	}
	return ""
}

// SetParams is...
func (c *Context) SetParams(params []param) {
	c.Params = params
}
