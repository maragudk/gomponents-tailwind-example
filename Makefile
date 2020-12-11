.PHONY: build-css start

build-css:
	NODE_ENV=production npx tailwindcss-cli@latest build views/app.css -o public/styles/app.css

start:
	go run cmd/server/*.go
