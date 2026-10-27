// Package health serves the HTTP health route probed by Kubernetes.
package health

import (
	"net/http"
)

const Path = "/health"

// Handler answers GET /health with a static OK body. The worker has no other
// HTTP surface; the route exists so the deployment's readiness and liveness
// probes have something to hit.
func Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(Path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			w.Header().Set("Allow", "GET, HEAD")
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})
	return mux
}

// ListenAndServe blocks serving Handler on addr.
func ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, Handler())
}
