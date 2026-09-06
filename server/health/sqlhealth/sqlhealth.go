package sqlhealth

import (
	"context"
	"database/sql"
)

type Checker struct {
	cancel context.CancelFunc

	stopped <-chan struct{}
	healthy bool
}

func New(db *sql.DB) *Checker { _ = "STUB: not implemented"; return nil }

func (c *Checker) CheckHealth() error { _ = "STUB: not implemented"; return nil }

func (c *Checker) Stop() { _ = "STUB: not implemented"; return }
