run: build
	@./bin/app

build: tmpl css
	@go build -o bin/app cmd/main.go

tmpl:
	@templ generate

css:
	tailwindcss -i views/css/app.css -o assets/styles.css --minify
