package middleware

import "net/http"

// SetHeadersHandler - middleware to set headers that are required for each request
type SetHeadersHandler struct {
	handler http.Handler
}

// NewSetHeadersHandler - return an instance of our SetHeadersHandler
func NewSetHeadersHandler(handler http.Handler) http.Handler {
	return &SetHeadersHandler{handler: handler}
}

func (h *SetHeadersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	skipHeadersRoutes := map[string]bool{
		"/GoogleLogin":    true,
		"/GoogleCallback": true,
	}
	if skipHeadersRoutes[path] == true {
		h.handler.ServeHTTP(w, r)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	h.handler.ServeHTTP(w, r)
}
