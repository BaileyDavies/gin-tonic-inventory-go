package middleware

import "github.com/julienschmidt/httprouter"

type Middleware func(handle httprouter.Handle) httprouter.Handle

// Accept mutiple functions of a middle ware type above and call them one by one to
// control the middleware flow

func Chain(f httprouter.Handle, m ...Middleware) httprouter.Handle {
	if len(m) == 0 {
		return f
	}
	return m[0](Chain(f, m[1:]...))
}
