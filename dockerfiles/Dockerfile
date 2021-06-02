FROM docker.io/nginx/unit:1.24.0-go1.15 AS build

ADD . /go/src/app
RUN cd /go/src/app && go build ./cmd/main.go

FROM docker.io/nginx/unit:1.24.0-go1.15

WORKDIR /go/src/app
COPY ./dockerfiles/config.json /docker-entrypoint.d/config.json
COPY --from=build /go/src/app/main .

EXPOSE 8080