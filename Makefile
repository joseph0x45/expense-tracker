BINARY=cashflow.out

build:
	cd frontend/ && npm run build
	go build -o $(BINARY) .

start-tailwind:
	npx tailwindcss -i ./public/input.css -o ./public/output.css --watch
