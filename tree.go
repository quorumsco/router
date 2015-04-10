package iogo

import "net/http"

type nodeType int

const (
	staticType nodeType = iota
	paramType
)

type node struct {
	nodeType  nodeType
	key       uint8
	handle    http.HandlerFunc
	param     string
	children  []*node
	name      string
	maxParams uint8
	path      string
	params    []param
}

type param struct {
	key   string
	value string
}

type NodeHeap []*node

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].key < h[j].key
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*node))
}

func (h *NodeHeap) Pop() interface{} {
	oldh := *h
	x := oldh[len(oldh)-1]
	newh := oldh[0 : len(oldh)-1]
	*h = newh
	return x
}

func newNode(key uint8, nodeType nodeType) *node {
	return &node{
		key:      key,
		children: make([]*node, 0),
		nodeType: nodeType,
	}
}

func findRec(n *node, path string, params []param) (http.HandlerFunc, bool, []param) {
	switch {
	case len(path) == 0:
		return n.handle, n.handle != nil, params
	default:
		for _, e := range n.children {
			if e.nodeType == staticType && e.key == path[0] {
				handle, found, par := findRec(e, path[1:], params)
				if found {
					return handle, found, par
				}
			}
		}
		for _, e := range n.children {
			if e.nodeType == paramType {
				p, param := consumeParameter(path)
				param.key = e.param
				handle, found, par := findRec(e, p, append(params, param))
				if found {
					return handle, found, par
				}
			}
		}
	}
	return nil, false, nil
}

func consumeParameter(path string) (string, param) {
	var i = 0
	for i < len(path) && path[i] != '/' {
		i++
	}
	return path[i:], param{
		value: path[:i],
	}
}

func (n *node) find(path string) (http.HandlerFunc, bool, []param) {
	return findRec(n, path, make([]param, 0, n.maxParams))
}

func insertRec(n *node, path string, handle http.HandlerFunc, params uint8) uint8 {
	switch {
	case len(path) == 0:
		n.handle = handle
		return params
	case path[0] == ':':
		p, param := consumeParameter(path[1:])
		for _, e := range n.children {
			if e.nodeType == paramType && e.param == param.value {
				return insertRec(e, p, handle, params+1)
			}
		}
		node := newNode(0, paramType)
		node.param = param.value
		n.children = append(n.children, node)
		return insertRec(node, p, handle, params+1)
	default:
		for _, e := range n.children {
			if e.nodeType == staticType && e.key == path[0] {
				return insertRec(e, path[1:], handle, params)
			}
		}
		node := newNode(path[0], staticType)
		n.children = append(n.children, node)
		return insertRec(node, path[1:], handle, params)
	}
}

func (n *node) insert(path string, handle http.HandlerFunc) {
	params := insertRec(n, path, handle, 0)
	if params > n.maxParams {
		n.maxParams = params
	}
}
