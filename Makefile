IMAGE := evns/crawler
APP := crawler

docker-build:
	docker build -t $(IMAGE) .

docker-run:
	docker run -ti $(IMAGE)

build:
	@echo "Building $(APP)"
	@go build -o $(APP)

run:
	@./$(APP)

test:
	go test -v

up: build run
