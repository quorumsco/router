package iogo

import "net/http"

type Router struct {
	table map[string]*node

	NotFound     http.HandlerFunc
	PanicHandler func(http.ResponseWriter, *http.Request, interface{})
}

func New() *Router {
	return &Router{
		table: make(map[string]*node),
	}
}

func (r *Router) Get(path string, handle http.HandlerFunc) {
	r.Handle("GET", path, handle)
}

func (r *Router) Post(path string, handle http.HandlerFunc) {
	r.Handle("POST", path, handle)
}

func (r *Router) Put(path string, handle http.HandlerFunc) {
	r.Handle("PUT", path, handle)
}

func (r *Router) Patch(path string, handle http.HandlerFunc) {
	r.Handle("PATCH", path, handle)
}

func (r *Router) Delete(path string, handle http.HandlerFunc) {
	r.Handle("DELETE", path, handle)
}

func (r *Router) Handle(method string, path string, handle http.HandlerFunc) {
	if path[0] != '/' {
		panic("can't handle relative path")
	}

	if r.table[method] == nil {
		r.table[method] = newNode(0, staticType)
	}

	r.table[method].insert(path, handle)
}

func (r *Router) recv(w http.ResponseWriter, req *http.Request) {
	if rcv := recover(); rcv != nil {
		r.PanicHandler(w, req, rcv)
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if r.PanicHandler != nil {
		defer r.recv(w, req)
	}

	if r.table[req.Method] != nil {
		path := req.URL.Path
		handle, found, params := r.table[req.Method].find(path)
		if found {
			ctx := setNewContext(req)
			ctx.SetParams(params)
			handle(w, req)
			return
		}
	}

	if r.NotFound != nil {
		r.NotFound(w, req)
	} else {
		http.NotFound(w, req)
	}
}
