FROM golang:1.18 as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server ./cmd


FROM scratch
COPY --from=builder /server /
EXPOSE 8000
ENTRYPOINT ["/server"]
