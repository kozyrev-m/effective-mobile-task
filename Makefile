test:
	go test -count=1 -v ./...

hello:
	@echo "Hello!"

.DEFAULT_GOAL = hello