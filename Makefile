run:
	@go run cmd/app/main.go

pub:
	@go run cmd/pub/main.go

test:
	@go test ./... -v

# initialize db at first
migrate-initial:
	cd migrations && goose postgres "user=admin password=pwd123 dbname=postgres sslmode=disable" up-by-one

# add all migrations to last
migrate-up:
	cd migrations && goose postgres "user=admin password=pwd123 dbname=orders sslmode=disable" up

# Roll back the version by 1
migrate-down:
	cd migrations && goose postgres "user=admin password=pwd123 dbname=orders sslmode=disable" down