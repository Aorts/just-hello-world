FROM golang:alpine as builder

WORKDIR /app
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o /just-hello-world

EXPOSE 8080

# Run
CMD ["/just-hello-world"]