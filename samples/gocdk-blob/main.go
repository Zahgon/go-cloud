package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"

	_ "gocloud.dev/blob/azureblob"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/s3blob"
)

const helpSuffix = `

  See https://gocloud.dev/concepts/urls/ for more background on
  Go CDK URLs, and sub-packages under gocloud.dev/blob
  (https://godoc.org/gocloud.dev/blob#pkg-subdirectories)
  for details on the blob.Bucket URL format.
`

func main() {
	os.Exit(run())
}

func run() int { _ = "STUB: not implemented"; return 0 }

type downloadCmd struct{}

func (*downloadCmd) Name() string     { _ = "STUB: not implemented"; return "" }
func (*downloadCmd) Synopsis() string { _ = "STUB: not implemented"; return "" }
func (*downloadCmd) Usage() string    { _ = "STUB: not implemented"; return "" }

func (*downloadCmd) SetFlags(_ *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (*downloadCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	_ = "STUB: not implemented"
	return *new(subcommands.ExitStatus)
}

type listCmd struct {
	prefix    string
	delimiter string
}

func (*listCmd) Name() string     { _ = "STUB: not implemented"; return "" }
func (*listCmd) Synopsis() string { _ = "STUB: not implemented"; return "" }
func (*listCmd) Usage() string    { _ = "STUB: not implemented"; return "" }

func (cmd *listCmd) SetFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (cmd *listCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	_ = "STUB: not implemented"
	return *new(subcommands.ExitStatus)
}

type uploadCmd struct{}

func (*uploadCmd) Name() string     { _ = "STUB: not implemented"; return "" }
func (*uploadCmd) Synopsis() string { _ = "STUB: not implemented"; return "" }
func (*uploadCmd) Usage() string    { _ = "STUB: not implemented"; return "" }

func (*uploadCmd) SetFlags(_ *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (*uploadCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...any) (status subcommands.ExitStatus) {
	_ = "STUB: not implemented"
	return *new(subcommands.ExitStatus)
}
