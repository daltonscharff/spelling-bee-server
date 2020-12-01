package updater

import (
	"database/sql"
	"fmt"

	// "github.com/daltonscharff/spelling-bee-server/utils/scraper/scraper"
	_ "github.com/lib/pq"
)

func Update(db *sql.DB) error {
	fmt.Println("Hello from update.go")
	return nil
}
