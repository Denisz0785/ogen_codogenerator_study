.PHONY: lint psql

lint:
	golangci-lint run ./...

psql:
	@echo "Connect to DB..."
	psql postgres://expense_user:zorro@localhost:5432/expenses