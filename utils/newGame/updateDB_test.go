package newGame

import (
	"testing"

	"github.com/daltonscharff/spelling-bee-server/config"
	"github.com/daltonscharff/spelling-bee-server/db"
)

var date string = "2020-12-01"
var letters []string = []string{"b", "i", "l", "m", "o", "t", "y"}
var center string = "t"
var words []string = []string{
	"blot",
	"blotto",
	"bolt",
	"boot",
	"booty",
	"bottom",
	"immobility",
	"itty",
	"lilt",
	"limit",
	"lobotomy",
	"loot",
	"lotto",
	"mitt",
	"mobility",
	"molt",
	"moot",
	"motility",
	"motto",
	"obit",
	"omit",
	"till",
	"tilt",
	"toil",
	"toll",
	"tomb",
	"tomboy",
	"tomtit",
	"tool",
	"toot",
}

func TestUpdate(t *testing.T) {
	conf, err := config.Read("../../config.yaml")
	if err != nil {
		panic(err)
	}
	db, err := db.Connect(conf)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	Update(db, date, letters, center, words)
}
