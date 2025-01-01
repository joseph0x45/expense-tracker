BINARY=cashflow-logger.out

build:
	go build -o $(BINARY) .

delete-db:
	rm $$HOME/.cashflow.db

test-direct-mode:
	@./$(BINARY) new -type=in -amount=100 -method='mobile money' \
		-planned=false -purpose='some lnog aaaah string'
