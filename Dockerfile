# build stage
FROM golang:alpine AS build
RUN apk update && apk add git
ADD . /src
RUN cd /src && go get -d ./... && go build -o crawler

# final stage
FROM alpine
WORKDIR /app
COPY --from=build /src/crawler /app/
ENTRYPOINT ./crawler
