package main

import "fmt"

type Middleware interface {
	Handle(request string, next Clouser)
}

type LoggingMiddleware struct{}

func (lm *LoggingMiddleware) Handle(request string, next Clouser) {
	fmt.Println("Logging request:", request)
	next(request)
}

type AuthenticationMiddleware struct{}

func (am *AuthenticationMiddleware) Handle(request string, next Clouser) {
	fmt.Println("Authenticating request:", request)
	next(request)
}

type ThrottlingMiddleware struct{}

func (tm *ThrottlingMiddleware) Handle(request string, next Clouser) {
	fmt.Println("Throttling request:", request)
	next(request)
}

type Kernel struct {
	middlewares []Middleware
}

func (k *Kernel) AddMiddleware(m Middleware) {
	k.middlewares = append(k.middlewares, m)
}

func (k *Kernel) Handle(request string) {
	// final handler
	var next Clouser = func(req string) {
		fmt.Println("Handling request:", req)
	}

	// build pipeline from last middleware to first
	for i := len(k.middlewares) - 1; i >= 0; i-- {
		mw := k.middlewares[i]
		prevNext := next
		next = func(req string) {
			mw.Handle(req, prevNext)
		}
	}

	// start the chain
	next(request)
}

type Clouser func(string)

func main() {
	kernel := &Kernel{}
	kernel.AddMiddleware(&LoggingMiddleware{})
	kernel.AddMiddleware(&AuthenticationMiddleware{})
	kernel.AddMiddleware(&ThrottlingMiddleware{})

	request := "GET /home"

	kernel.Handle(request)
}
