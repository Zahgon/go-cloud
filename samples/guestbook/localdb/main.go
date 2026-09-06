package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	guestbookDir := flag.String("guestbook_dir", ".", "directory containing guestbook sample source code")
	flag.Parse()
	if flag.NArg() > 1 {
		fmt.Fprintf(os.Stderr, "usage: localdb [flags] container_name\n")
		os.Exit(1)
	}
	log.SetPrefix("localdb: ")
	log.SetFlags(0)
	if err := runLocalDB(flag.Arg(0), *guestbookDir); err != nil {
		log.Fatal(err)
	}
}

func runLocalDB(containerName, guestbookDir string) error { _ = "STUB: not implemented"; return nil }
