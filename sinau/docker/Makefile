swag:
	@swag init --dir presenter --generalInfo ../cmd/main.go --output presenter/handler/docs

docker:
	@docker build ./ -t sinau_app:1.0

docker-compose:
	@docker-compose up --build

build-linux:
	@GOOS=linux GOARCH=amd64 go build -o bin/sinau-app ./cmd

build-win:
	@GOOS=windows GOARCH=amd64 go build -o bin/sinau-app.exe ./cmd

build-mac:
	@GOOS=darwin GOARCH=amd64 go build -o bin/sinau-app ./cmd

