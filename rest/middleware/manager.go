package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(
		mngr.globalMiddlewares,
		middlewares...,
	)
}

func (mngr *Manager) GetGlobalMiddlewares() []Middleware {
	return mngr.globalMiddlewares
}

func (mngr *Manager) With(middlewares ...Middleware) Middleware {

	return func(next http.Handler) http.Handler {

		n := next

		for i := len(middlewares) - 1; i >= 0; i-- {
			n = middlewares[i](n)
		}

		return n
	}
}

func (mngr *Manager) Handle(
	mux *http.ServeMux,
	path string,
	handler http.HandlerFunc,
	middlewares ...Middleware,
) {

	wrappedHandler := mngr.With(middlewares...)(handler)
	mux.Handle(path, wrappedHandler)
}
