BINARY=cashflow.out

build:
	go build -o $(BINARY) .

start-tailwind:
	npx tailwindcss -i ./public/input.css -o ./public/output.css --watch
