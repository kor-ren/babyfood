build: build-app build-go

ci: restore-go restore-npm build

generate:
	echo "Generate GraphQL Model"
	go run github.com/99designs/gqlgen generate

	echo "Generate Apollo"
	cd ./app && npm run compile


restore-npm:
	cd ./app && npm ci

restore-go:
	cd ./server && go mod download

build-go:
	echo "Build server"
	cd ./server && CGO_ENABLED=0 GOOS=linux go build -o ../babyfood-linux

build-app:
	echo "Build app"
	cd ./app && npm run build


build-docker: build
	docker build -t test .