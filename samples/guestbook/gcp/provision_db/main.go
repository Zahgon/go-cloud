package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("gcp/provision_db: ")
	project := flag.String("project", "", "GCP project ID")
	serviceAccount := flag.String("service_account", "", "name of service account in GCP project")
	instance := flag.String("instance", "", "database instance name")
	database := flag.String("database", "", "name of database to initialize")
	password := flag.String("password", "", "root password for the database")
	schema := flag.String("schema", "", "path to .sql file defining the database schema")
	flag.Parse()
	missing := false
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() == "" {
			log.Printf("Required flag -%s is not set.", f.Name)
			missing = true
		}
	})
	if missing {
		os.Exit(64)
	}
	if err := provisionDB(*project, *serviceAccount, *instance, *database, *password, *schema); err != nil {
		log.Fatal(err)
	}
}

type key struct {
	PrivateKeyID string `json:"private_key_id"`
}

func provisionDB(projectID, serviceAccount, dbInstance, dbName, dbPassword, schemaPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func run(args ...string) (stdout string, err error) { _ = "STUB: not implemented"; return "", nil }

type gcloud struct {
	project string
}

func (gcp *gcloud) cmd(args ...string) []string { _ = "STUB: not implemented"; return nil }
