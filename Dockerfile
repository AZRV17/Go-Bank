FROM golang:1.21-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg
RUN go build -o ./bin/app cmd/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/app /
COPY internal/config/config.yaml internal/config/config.yaml

EXPOSE 8080
CMD [ "/app" ]
