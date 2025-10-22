package main

//go:generate bash scripts/migrate.sh
//go:generate sqlc generate -f ./internal/pgstore/sqlc.yaml
