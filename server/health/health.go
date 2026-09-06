package health

import (
	"net/http"
)

type Handler struct {
	checkers []Checker
}

func (h *Handler) Add(c Checker) { _ = "STUB: not implemented"; return }

func (h *Handler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

func writeHeaders(statusLen string, w http.ResponseWriter) { _ = "STUB: not implemented"; return }

func writeUnhealthy(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

func HandleLive(w http.ResponseWriter, _ *http.Request) { _ = "STUB: not implemented"; return }

func writeHealthy(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

type Checker interface {
	CheckHealth() error
}

type CheckerFunc func() error

func (f CheckerFunc) CheckHealth() error { _ = "STUB: not implemented"; return nil }
