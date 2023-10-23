package pg

import (
	"log"
	"strings"

	"github.com/integralist/go-findroot/find"
)

// defaultMigrationPath is a default path to migrations.
var defaultMigrationPath = "/migrations"

// getMigrationPath gets migration path.
func getMigrationPath() string {
	// get root dir
	rep, err := find.Repo()
	if err != nil {
		log.Printf("cannot get root dir: %s", err.Error())
		log.Println("use default path to migration")
		return defaultMigrationPath
	}

	path := strings.Join([]string{rep.Path, "migrations"}, "/")

	if strings.EqualFold(rep.Path, "./") || strings.EqualFold(rep.Path, "/") {
		path = strings.Join([]string{rep.Path, "migrations"}, "")
	}

	return path
}
