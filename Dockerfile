FROM golang:1.13.5 AS builder
WORKDIR /go/src/satoshiyamamoto/go-gin-gorm-todos/
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/satoshiyamamoto/go-gin-gorm-todos/app .
CMD ["./app"]