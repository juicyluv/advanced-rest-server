.SILENT:
.PHONY:

run:
	go run cmd/api/main.go

run-dev:
	go run cmd/api/main.go -config-name server-dev

migrate-up:
	migrate -path=migrations -database="postgres://audiosystem:pa55word@localhost:5432/audiosystem?sslmode=disable" up

migrate-down:
	migrate -path=migrations -database="postgres://audiosystem:pa55word@localhost:5432/audiosystem?sslmode=disable" down