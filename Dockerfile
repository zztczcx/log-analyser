FROM golang:1.21-alpine3.18 as builder

WORKDIR /build

COPY . /build/log-analyser

WORKDIR /build/log-analyser

RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 go build -a --ldflags "-s -w" -o /build/main ./

FROM alpine:3.18

# Set workdir on current image
WORKDIR /app
# Leverage a separate non-root user for the application
RUN adduser -S -D -H -h /app mantel
# Change to a non-root user
USER mantel
# Add artifact from builder stage
COPY --from=builder /build/main /app/

ENTRYPOINT ["./main"]

CMD ["-h"]
