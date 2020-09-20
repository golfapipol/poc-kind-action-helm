FROM golang:1.14 as builder
ARG GIT_COMMIT
WORKDIR /module
COPY ./go.mod /module/go.mod
COPY ./go.sum /module/go.sum
RUN go mod download
COPY . /module
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.GitCommit=$GIT_COMMIT" -o ./bin/app ./cmd/main.go

FROM alpine:3
RUN apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime && \
    echo "Asia/Bangkok" >  /etc/timezone && \
    apk del tzdata
WORKDIR /root/
COPY --from=builder /module/bin .
ENV GIN_MODE release
EXPOSE 8080
CMD ./app