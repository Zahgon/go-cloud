package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("aws/provision_db: ")
	host := flag.String("host", "", "hostname of database")
	region := flag.String("region", "", "AWS region")
	securityGroup := flag.String("security_group", "", "database security group")
	database := flag.String("database", "", "name of database to provision")
	password := flag.String("password", "", "root password on database")
	schema := flag.String("schema", "", "path to .sql file defining the database schema")
	flag.Parse()
	missing := false
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() == "" {
			log.Printf("Required flag -%s not set.", f.Name)
			missing = true
		}
	})
	if missing {
		os.Exit(64)
	}
	if err := provisionDb(*host, *region, *securityGroup, *database, *password, *schema); err != nil {
		log.Fatal(err)
	}
}

func provisionDb(dbHost, region, securityGroupID, dbName, dbPassword, schemaPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func run(args ...string) (stdout string, err error) { _ = "STUB: not implemented"; return "", nil }
