.PHONY: run
run:
				go run ./cmd/server/main.go

.PHONY: build
build:
				go build ./cmd/server/main.go

.PHONY: migrateup
migrateup:
			migrate -path migrations -database "postgres://localhost/invest?user=raduga&password=faster123" up
			
.PHONY: migratedown
migratedown:
			migrate -path migrations -database "postgres://localhost/invest?user=raduga&password=faster123" down

.PHONY: migrateforce
migrateforce:
			migrate -path migrations -database "postgres://localhost/invest?user=raduga&password=faster123" force $(version)

.PHONY: migratecreate
migratecreate:
			migrate create -ext sql -dir migrations $(name)




.DEFAULT_GOAL := run