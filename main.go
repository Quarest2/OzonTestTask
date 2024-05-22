package main

import (
	"OzonTestTask/OzonTestTask/db"
	"os"
)

var MemoryFlag = 0

func main() {
	dbOrMemory, ok := os.LookupEnv("MEMORY")
	if !ok || dbOrMemory != "MEMORY" {
		db.Connect()
	} else {
		MemoryFlag = 1
	}
}
