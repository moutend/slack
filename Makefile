prepare:
	@go get -u -t github.com/volatiletech/sqlboiler
	@go get -u -t github.com/volatiletech/sqlboiler-sqlite3

generate:
	@echo "generate models ..."
	@go run _tools/up.go
	@sqlboiler sqlite3 --output ./internal/models
	@rm temporary.db3
