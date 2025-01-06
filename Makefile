BINARY=cashflow.out

build:
	npx tailwindcss -i ./web/public/input.css -o ./web/public/output.css
	go build -o $(BINARY) .

start-tailwind:
	npx tailwindcss -i ./web/public/input.css -o ./web/public/output.css --watch
