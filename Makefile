.PHONY: build start watch

build:
	NODE_ENV=production ./node_modules/.bin/postcss build tailwind.css -o public/styles/app.css

start:
	go run cmd/server/*.go

watch:
	NODE_ENV=development ./node_modules/.bin/postcss build tailwind.css -o public/styles/app.css -w
