BINARY=cashflow.out

build:
	npx tailwindcss -i ./public/input.css -o ./public/output.css
	go build -o $(BINARY) .

compile-tailwind:
	npx tailwindcss -i ./public/input.css -o ./public/output.css

start-tailwind:
	npx tailwindcss -i ./public/input.css -o ./public/output.css --watch
