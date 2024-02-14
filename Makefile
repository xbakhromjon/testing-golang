include .env
export

test:
	@echo "Loading env"
	export DATABASE_URL=${DATABASE_URL}
	@echo "Running tests"
	cd "./internal"
	go test -count=1 -p=8 -parallel=8 -race ./...
