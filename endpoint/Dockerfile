FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

# Fetching dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# Build with some optimization
RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

EXPOSE 8000

CMD ["/dist/main"]