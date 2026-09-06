package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"

	_ "gocloud.dev/secrets/awskms"
	_ "gocloud.dev/secrets/azurekeyvault"
	_ "gocloud.dev/secrets/gcpkms"
	_ "gocloud.dev/secrets/hashivault"
	_ "gocloud.dev/secrets/localsecrets"
)

const helpSuffix = `

  See https://gocloud.dev/concepts/urls/ for more background on
  Go CDK URLs, and sub-packages under gocloud.dev/secrets
  (https://godoc.org/gocloud.dev/secrets#pkg-subdirectories)
  for details on the secrets.Keeper URL format.
`

func main() {
	os.Exit(run())
}

func run() int { _ = "STUB: not implemented"; return 0 }

type decryptCmd struct {
	base64in  bool
	base64out bool
}

func (*decryptCmd) Name() string     { _ = "STUB: not implemented"; return "" }
func (*decryptCmd) Synopsis() string { _ = "STUB: not implemented"; return "" }
func (*decryptCmd) Usage() string    { _ = "STUB: not implemented"; return "" }

func (cmd *decryptCmd) SetFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (cmd *decryptCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	_ = "STUB: not implemented"
	return *new(subcommands.ExitStatus)
}

type encryptCmd struct {
	base64in  bool
	base64out bool
}

func (*encryptCmd) Name() string     { _ = "STUB: not implemented"; return "" }
func (*encryptCmd) Synopsis() string { _ = "STUB: not implemented"; return "" }
func (*encryptCmd) Usage() string    { _ = "STUB: not implemented"; return "" }

func (cmd *encryptCmd) SetFlags(f *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (cmd *encryptCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	_ = "STUB: not implemented"
	return *new(subcommands.ExitStatus)
}
