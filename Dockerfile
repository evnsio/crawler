# build stage
FROM golang:alpine AS build
RUN apk update && apk add git
ADD . /src
RUN cd /src && go get -d ./... && go build -o crawler

# final stage
FROM alpine
RUN apk --update upgrade && \
    apk add curl ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*
WORKDIR /
COPY --from=build /src/crawler /
CMD ["/crawler"]
