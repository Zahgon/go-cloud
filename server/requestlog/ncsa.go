package requestlog

import (
	"io"
	"sync"
)

type NCSALogger struct {
	onErr func(error)

	mu  sync.Mutex
	w   io.Writer
	buf []byte
}

func NewNCSALogger(w io.Writer, onErr func(error)) *NCSALogger {
	_ = "STUB: not implemented"
	return nil
}

func (l *NCSALogger) Log(ent *Entry) { _ = "STUB: not implemented"; return }

func (l *NCSALogger) log(ent *Entry) error { _ = "STUB: not implemented"; return nil }

func formatEntry(b []byte, ent *Entry) []byte { _ = "STUB: not implemented"; return nil }
