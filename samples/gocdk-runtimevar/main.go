package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"

	_ "gocloud.dev/runtimevar/awsparamstore"
	_ "gocloud.dev/runtimevar/blobvar"
	_ "gocloud.dev/runtimevar/constantvar"
	_ "gocloud.dev/runtimevar/filevar"
	_ "gocloud.dev/runtimevar/gcpruntimeconfig"
	_ "gocloud.dev/runtimevar/hashivault"
	_ "gocloud.dev/runtimevar/httpvar"
)

const helpSuffix = `

  See https://gocloud.dev/concepts/urls/ for more background on
  Go CDK URLs, and sub-packages under gocloud.dev/runtimevar
  (https://godoc.org/gocloud.dev/runtimevar#pkg-subdirectories)
  for details on the runtimevar.Variable URL format.
`

func main() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int { _ = "STUB: not implemented"; return 0 }

type catCmd struct{}

func (*catCmd) Name() string     { _ = "STUB: not implemented"; return "" }
func (*catCmd) Synopsis() string { _ = "STUB: not implemented"; return "" }
func (*catCmd) Usage() string    { _ = "STUB: not implemented"; return "" }

func (*catCmd) SetFlags(_ *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (*catCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	_ = "STUB: not implemented"
	return *new(subcommands.ExitStatus)
}

type watchCmd struct{}

func (*watchCmd) Name() string     { _ = "STUB: not implemented"; return "" }
func (*watchCmd) Synopsis() string { _ = "STUB: not implemented"; return "" }
func (*watchCmd) Usage() string    { _ = "STUB: not implemented"; return "" }

func (*watchCmd) SetFlags(_ *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (*watchCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	_ = "STUB: not implemented"
	return *new(subcommands.ExitStatus)
}
