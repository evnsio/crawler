IMAGE := evns/crawler

build:
	docker build -t $(IMAGE) .

run:
	docker run -ti -p 8080:8080 $(IMAGE)

up: build run
