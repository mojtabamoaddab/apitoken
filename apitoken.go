package apitoken

import (
	"net/http"
)

type Handler struct {
	Handler             http.Handler
	UnauthorizedHandler http.Handler
	TokenSet            TokenSet
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-API-Token")

	if !h.TokenSet.Validate(token) {
		h.UnauthorizedHandler.ServeHTTP(w, r)
		return
	}

	h.Handler.ServeHTTP(w, r)
}

func NewHandler(h http.Handler, ts TokenSet) *Handler {
	return &Handler{
		Handler:             h,
		TokenSet:            ts,
		UnauthorizedHandler: DefaultUnautorizedHandler,
	}
}

func defaultUnautorizedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("API authentication failed.\n"))
}

var DefaultUnautorizedHandler = http.HandlerFunc(defaultUnautorizedHandler)
