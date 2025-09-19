run:
	go run cmd/main.go

compile: tailwindcss templ run 

vite: 
	bun run vite ./public 

watch:
	find . -name "*.templ" | entr -r make compile

tailwindcss:
	bun run tailwindcss --config tailwind.config.js

templ:
	templ generate ./components

.PHONY: run compile vite watch test tailwindcss templ