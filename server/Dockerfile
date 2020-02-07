FROM golang:1.13.4 as builder
WORKDIR /golang
COPY . .
# Set Environment Variable
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
# Build
RUN go build -o app main.go

# Runtime Container
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /golang/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
