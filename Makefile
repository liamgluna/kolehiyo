run:
	go run ./cmd/api

psql:
	@psql ${KOLEHIYO_DB_DSN}

migration:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

migrate-up:
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${KOLEHIYO_DB_DSN} up

migrate-down:
	@echo 'Running down migrations...'
	migrate -path ./migrations -database ${KOLEHIYO_DB_DSN} down