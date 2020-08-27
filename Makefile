generate:
	@echo "generate models ..."
	go run _tools/up.go
	sqlboiler sqlite3 --output ./internal/models
	rm temporary.db3
