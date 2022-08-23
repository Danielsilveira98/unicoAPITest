FROM golang:1.18 as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server ./cmd/api


FROM scratch
COPY --from=builder /server /
COPY --from=builder /app/deployment/migrations /
# run ls
EXPOSE 8000
ENTRYPOINT ["/server"]
