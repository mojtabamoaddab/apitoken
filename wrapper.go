package apitoken

import (
	"net/http"
)

type Wrapper struct {
	UnauthorizedHandler http.Handler
	TokenSet            TokenSet
}

func (w *Wrapper) Wrap(h http.Handler) http.Handler {
	return &Handler{
		Handler:             h,
		UnauthorizedHandler: w.UnauthorizedHandler,
		TokenSet:            w.TokenSet,
	}
}

func (w *Wrapper) WrapFunc(f func(http.ResponseWriter, *http.Request)) http.Handler {
	return w.Wrap(http.HandlerFunc(f))
}

func NewWrapper(tokens TokenSet) *Wrapper {
	return &Wrapper{
		UnauthorizedHandler: DefaultUnautorizedHandler,
		TokenSet:            tokens,
	}
}
