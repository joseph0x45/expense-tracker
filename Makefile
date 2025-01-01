BINARY=cashflow-logger.out

build:
	go build -o $(BINARY) .

delete-db:
	rm $$HOME/.cashflow.db

test-cash-in:
	@./$(BINARY) new -type=in -amount=100 -method='mobile money' \
		-planned=false -purpose='some lnog aaaah string'

test-cash-out:
	@./$(BINARY) new -type=out -amount=100 -method='mobile money' \
		-planned=false -purpose='some lnog aaaah string'

get-balance:
	@./$(BINARY) get-current-balance
