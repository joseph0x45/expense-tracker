BINARY=cashflow-logger.out

build:
	go build -o $(BINARY) .

delete-db:
	rm $$HOME/.cashflow.db
