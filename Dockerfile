#Build binary
FROM golang:alpine as builder
LABEL maintainer = "Mohammed Osama"
RUN apk update && apk add --no-cache git
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENDABLED=0 GOOS=linux go build -a -installsufix cgo -o main

# Builder container to run binary
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/config config
EXPOSE 8000
ENTRYPOINT ["./main"]