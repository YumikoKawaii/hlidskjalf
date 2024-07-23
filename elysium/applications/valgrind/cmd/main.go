package main

import (
	"elysium.com/applications/valgrind/config"
	"elysium.com/applications/valgrind/serve"
	"fmt"
)

func main() {

	cfg, kongCtx := config.Initialize()
	switch kongCtx.Command() {
	case "migrate-schema":
		fmt.Println("migrating schema...")
		serve.Migrate(cfg)
		return
		//default:
		//	log.Fatalf("unexpected command: %v", kongCtx.Command())
	}

}
