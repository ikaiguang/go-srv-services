package main

import (
	dbmigrate "github.com/ikaiguang/go-srv-services/app/snowflake-service/cmd/migration/migrate"
)

// go run ./cmd/migration/... -conf=./configs
func main() {
	dbmigrate.Run(dbmigrate.WithCloseEngineHandler())
}
