package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"

	_ "gocloud.dev/pubsub/awssnssqs"
	_ "gocloud.dev/pubsub/azuresb"
	_ "gocloud.dev/pubsub/gcppubsub"
	_ "gocloud.dev/pubsub/kafkapubsub"
	_ "gocloud.dev/pubsub/natspubsub"
	_ "gocloud.dev/pubsub/rabbitpubsub"
)

const helpSuffix = `

  See https://gocloud.dev/concepts/urls/ for more background on
  Go CDK URLs, and sub-packages under gocloud.dev/pubsub
  (https://godoc.org/gocloud.dev/pubsub#pkg-subdirectories)
  for details on the topic/subscription URL format.
`

func main() {
	os.Exit(run())
}

func run() int { _ = "STUB: not implemented"; return 0 }

type pubCmd struct{}

func (*pubCmd) Name() string     { _ = "STUB: not implemented"; return "" }
func (*pubCmd) Synopsis() string { _ = "STUB: not implemented"; return "" }
func (*pubCmd) Usage() string    { _ = "STUB: not implemented"; return "" }

func (*pubCmd) SetFlags(_ *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (*pubCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	_ = "STUB: not implemented"
	return *new(subcommands.ExitStatus)
}

type subCmd struct {
	n int
}

func (*subCmd) Name() string     { _ = "STUB: not implemented"; return "" }
func (*subCmd) Synopsis() string { _ = "STUB: not implemented"; return "" }
func (*subCmd) Usage() string    { _ = "STUB: not implemented"; return "" }

func (cmd *subCmd) SetFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (cmd *subCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	_ = "STUB: not implemented"
	return *new(subcommands.ExitStatus)
}
