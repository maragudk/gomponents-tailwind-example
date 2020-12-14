.PHONY: build-css build-css-dev start

build-css:
	NODE_ENV=production npx tailwindcss-cli@latest build views/app.css -o public/styles/app.css

build-css-dev:
	NODE_ENV=development npx tailwindcss-cli@latest build views/app.css -o public/styles/app.css

start:
	go run cmd/server/*.go
