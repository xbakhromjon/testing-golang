
test:
	@echo "Running tests"
	cd "./internal"
	go test -count=1 -p=8 -parallel=8 -race ./...
