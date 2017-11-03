IMAGE := evns/crawler
APP := crawler

docker-build:
	docker build -t $(IMAGE) .

build:
	@echo "Building $(APP)"
	@go build -o $(APP)
	@echo "Done"

run:
	@./$(APP) ${ARGS}

test:
	@echo "Running Tests"
	@go test -v
	@echo "Done"

go: build run
