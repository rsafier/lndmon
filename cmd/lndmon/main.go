package main

import (
	"github.com/rsafier/lndmon"

	_ "github.com/rsafier/lndmon/collectors"
)

func main() {
	lndmon.Main()
}
