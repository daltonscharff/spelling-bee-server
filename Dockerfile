FROM golang:1.15-alpine
WORKDIR /go/src

COPY go.* .
RUN go mod download
COPY . .

CMD [ "go", "run", "/go/src/cmd/spelling-bee/main.go" ]