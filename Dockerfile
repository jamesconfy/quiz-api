# #build stage
# FROM golang:alpine AS builder
# RUN apk add --no-cache git
# WORKDIR /go/src/app
# COPY . .
# RUN go get -d -v ./...
# RUN go build -o /go/bin/app -v ./...

# #final stage
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# COPY --from=builder /go/bin/app /app
# ENTRYPOINT /app
# LABEL Name=quizapi Version=0.0.1
# EXPOSE 8000
# syntax=docker/dockerfile:1

FROM golang:1.19-alpine3.16

RUN mkdir /quiz-api

COPY . /quiz-api

WORKDIR /quiz-api

RUN go build -o main

EXPOSE  8080

CMD [ "./main" ]