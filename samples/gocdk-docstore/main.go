package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"

	_ "gocloud.dev/docstore/awsdynamodb/v2"
	_ "gocloud.dev/docstore/gcpfirestore"
	_ "gocloud.dev/docstore/memdocstore"
	_ "gocloud.dev/docstore/mongodocstore"
)

const helpSuffix = `

  See https://gocloud.dev/concepts/urls/ for more background on
  Go CDK URLs, and sub-packages under gocloud.dev/docstore
  (https://godoc.org/gocloud.dev/docstore#pkg-subdirectories)
  for details on the docstore.Collection URL format.
`

func main() {
	os.Exit(run())
}

func run() int { _ = "STUB: not implemented"; return 0 }

type Message struct {
	ID               string
	Date             string
	Content          string
	DocstoreRevision any
}

func (m Message) String() string { _ = "STUB: not implemented"; return "" }

type listCmd struct {
	date string
}

func (*listCmd) Name() string     { _ = "STUB: not implemented"; return "" }
func (*listCmd) Synopsis() string { _ = "STUB: not implemented"; return "" }
func (*listCmd) Usage() string    { _ = "STUB: not implemented"; return "" }

func (cmd *listCmd) SetFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (cmd *listCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	_ = "STUB: not implemented"
	return *new(subcommands.ExitStatus)
}

type putCmd struct {
	id   string
	date string
}

func (*putCmd) Name() string     { _ = "STUB: not implemented"; return "" }
func (*putCmd) Synopsis() string { _ = "STUB: not implemented"; return "" }
func (*putCmd) Usage() string    { _ = "STUB: not implemented"; return "" }

func (p *putCmd) SetFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (p *putCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	_ = "STUB: not implemented"
	return *new(subcommands.ExitStatus)
}

type updateCmd struct{}

func (*updateCmd) Name() string     { _ = "STUB: not implemented"; return "" }
func (*updateCmd) Synopsis() string { _ = "STUB: not implemented"; return "" }
func (*updateCmd) Usage() string    { _ = "STUB: not implemented"; return "" }

func (*updateCmd) SetFlags(_ *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (cmd *updateCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	_ = "STUB: not implemented"
	return *new(subcommands.ExitStatus)
}

type deleteCmd struct {
	date string
}

func (*deleteCmd) Name() string     { _ = "STUB: not implemented"; return "" }
func (*deleteCmd) Synopsis() string { _ = "STUB: not implemented"; return "" }
func (*deleteCmd) Usage() string    { _ = "STUB: not implemented"; return "" }

func (cmd *deleteCmd) SetFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (cmd *deleteCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	_ = "STUB: not implemented"
	return *new(subcommands.ExitStatus)
}
