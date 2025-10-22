#!/usr/bin/env bash
set -a
source .env
set +a

chmod +x scripts/migrate.sh

tern migrate --migrations ./internal/pgstore/migrations --config ./internal/pgstore/migrations/tern.conf
