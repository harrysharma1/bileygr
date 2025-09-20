run:
	go run cmd/main.go

compile: tailwindcss templ run 

vite: 
	bun run vite ./public 

watch:
	find . -name "*.templ" | entr -r make compile

tailwindcss:
	bun run tailwindcss --config tailwind.config.js -i ./static/css/input.css -o ./static/css/output.css \
		--content "./components/**/*.templ"

templ:
	templ generate --watch --proxy="http://localhost:6969" --cmd="go run cmd/main.go" ./components

.PHONY: run compile vite watch test tailwindcss templ