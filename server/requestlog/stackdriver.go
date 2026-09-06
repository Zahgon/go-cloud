package requestlog

import (
	"bytes"
	"encoding/json"
	"io"
	"sync"
	"time"
)

type StackdriverLogger struct {
	onErr func(error)

	mu  sync.Mutex
	w   io.Writer
	buf bytes.Buffer
	enc *json.Encoder
}

func NewStackdriverLogger(w io.Writer, onErr func(error)) *StackdriverLogger {
	_ = "STUB: not implemented"
	return nil
}

func (l *StackdriverLogger) Log(ent *Entry) { _ = "STUB: not implemented"; return }

func (l *StackdriverLogger) log(ent *Entry) error { _ = "STUB: not implemented"; return nil }

func appendLatency(b []byte, d time.Duration) []byte { _ = "STUB: not implemented"; return nil }
