package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Chain struct {
	middlewares []Middleware
}

func NewChain(middlewares ...Middleware) *Chain {
	return &Chain{middlewares: middlewares}
}

func (c *Chain) Then(finalHandler http.Handler) http.Handler {
	if finalHandler == nil {
		finalHandler = http.DefaultServeMux
	}

	// Apply from the back to the front
	for i := len(c.middlewares) - 1; i >= 0; i-- {
		finalHandler = c.middlewares[i](finalHandler)
	}
	return finalHandler
}
