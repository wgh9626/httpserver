FROM golang:1.17.5 AS build-env

ADD . /app
WORKDIR /app

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/http_server

FROM alpine AS final

COPY --from=build-env /app/http_server /

EXPOSE 8080
CMD [ "/http_server" ]
