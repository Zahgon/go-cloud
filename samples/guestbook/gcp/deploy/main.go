package main

import (
	"flag"
	"log"
	"os/exec"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("gcp/deploy: ")
	guestbookDir := flag.String("guestbook_dir", "..", "directory containing the guestbook example")
	tfStatePath := flag.String("tfstate", "terraform.tfstate", "path to terraform state file")
	flag.Parse()
	if err := deploy(*guestbookDir, *tfStatePath); err != nil {
		log.Fatal(err)
	}
}

func deploy(guestbookDir, tfStatePath string) error { _ = "STUB: not implemented"; return nil }

type (
	service      struct{ Status *status }
	status       struct{ LoadBalancer loadBalancer }
	loadBalancer struct{ Ingress []ingress }
	ingress      struct{ IP string }
)

type gcloud struct {
	projectID string
}

func (gcp *gcloud) cmd(args ...string) *exec.Cmd { _ = "STUB: not implemented"; return nil }

func run(args ...string) (stdout string, err error) { _ = "STUB: not implemented"; return "", nil }

func runb(args ...string) (stdout []byte, err error) { _ = "STUB: not implemented"; return nil, nil }
