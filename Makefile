.SILENT:
.PHONY:

run:
	go run cmd/api/main.go

run-dev:
	go run cmd/api/main.go -config-name server-dev