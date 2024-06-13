build: build-app

ci: restore-npm build

generate:
	echo "Generate GraphQL Model"
	cd ./server && go run github.com/99designs/gqlgen generate

	echo "Generate Apollo"
	cd ./app && npm run compile


restore-npm:
	cd ./app && npm ci

build-app:
	echo "Build app"
	cd ./app && npm run build


build-docker: build
	docker build -t test .