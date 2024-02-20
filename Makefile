include .envrc

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	go run ./cmd/api -db-dsn=${KOLEHIYO_DB_DSN}

## run/api: run the cmd/api application and enalbe CO
.PHONY: run/api/cors
run/api/cors:
	go run ./cmd/api -db-dsn=${KOLEHIYO_DB_DSN} -cors-trusted-origins=${name}

## db/psql: connect to the database using psql
.PHONY: db/psql
db/psql:
	@psql ${KOLEHIYO_DB_DSN}

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up confirm
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${KOLEHIYO_DB_DSN} up

## db/migrations/down: apply all down database migrations
.PHONY: db/migrations/down confirm
db/migrations/down: confirm
	@echo 'Running down migrations...'
	migrate -path ./migrations -database ${KOLEHIYO_DB_DSN} down

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy and vendor dependencies and format, vet and test all code
.PHONY: audit
audit: vendor
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:	
	@echo 'Tidying and veryfing module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/api: build the cmd/api application
.PHONY: build/api
build/api:
	@echo 'Building cmd/api...'
	go build -ldflags='-s' -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./bin/linux_amd64/api ./cmd/api

# ==================================================================================== #
# PRODUCTION
# ==================================================================================== #

production_host_ip = '128.199.172.210'

## production/connect: connect to the production server
.PHONY: production/connect
production/connect:
	ssh kolehiyo@${production_host_ip}

## production/deploy/api: deploy the api to production
.PHONY: production/deploy/api
production/deploy/api:
	rsync -P ./bin/linux_amd64/api kolehiyo@${production_host_ip}:~
	rsync -rP --delete ./migrations kolehiyo@${production_host_ip}:~
	rsync -P ./remote/production/api.service kolehiyo@${production_host_ip}:~
	rsync -P ./remote/production/Caddyfile kolehiyo@${production_host_ip}:~
	ssh -t kolehiyo@${production_host_ip} '\
		migrate -path ~/migrations -database $$KOLEHIYO_DB_DSN up \
		&& sudo mv ~/api.service /etc/systemd/system/ \
		&& sudo systemctl enable api \
		&& sudo systemctl restart api \
		&& sudo mv ~/Caddyfile /etc/caddy/ \
		&& sudo systemctl reload caddy \
	'