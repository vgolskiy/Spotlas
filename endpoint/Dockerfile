FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app
COPY . .

# Fetching dependencies
RUN go get -d -v

RUN go build -o main

# Build final image
FROM scratch AS final

WORKDIR /app

COPY --from=builder /app/main /app

EXPOSE 8000

CMD ["/app/main"]