package util

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

// A function that help chain multiple middlewares to one route in right to left order
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
