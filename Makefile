BINARY=expense_tracker.out
DB_FILE=/home/joseph/.expense_tracker.db

build:
	go build -o $(BINARY) .

rm-db:
	rm $(DB_FILE)

view-db:
	cat $(DB_FILE)

into-db:
	sqlite3 $(DB_FILE)

tailwind-compile:
	npx @tailwindcss/cli -i ./input.css -o ./output.css

tailwind-watch:
	npx @tailwindcss/cli -i ./input.css -o ./output.css --watch
