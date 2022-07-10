
FROM golang:1.18-alpine AS builder

WORKDIR /app

# Copy (go.mod, go.sum) and install dependencies in /app
COPY /API/go.mod /API/go.sum ./
RUN go mod download

# copy sourse code to /app
COPY . .

WORKDIR /app/API

# Build app
RUN go build -o main .

# ---------- multi-stage --------------
FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/API/main /app/API/main

EXPOSE 8081

# Run the binary program produced by `go install`
CMD ["/app/API/main"]
