package handlers

import "net/http"

func (m *Mines) MiddleWarePullGame(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}
