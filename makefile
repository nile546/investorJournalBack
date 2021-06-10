.PHONY: run
run:
				go run ./cmd/server/main.go

.PHONY: build
build:
				go build ./cmd/server/main.go

.PHONY: migrateup
migrateup:
			migrate -path migrations -database "postgres://localhost/invest?user=invest&password=fast" up
			
.PHONY: migratedown
migratedown:
			migrate -path migrations -database "postgres://localhost/invest?user=invest&password=fast" down

.PHONY: migrateforce
migrateforce:
			migrate -path migrations -database "postgres://localhost/invest?user=invest&password=fast" force $(version)

.PHONY: migratecreate
migratecreate:
			migrate create -ext sql -dir migrations $(name)




.DEFAULT_GOAL := run