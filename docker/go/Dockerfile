FROM golang:1.18-alpine as builder

WORKDIR /app

COPY ./go.mod .
RUN go mod download

COPY . .

RUN go build -o /bin/app ./cmd

FROM builder as dev

COPY --from=builder /bin/app /bin/app

ENTRYPOINT ["/bin/app"]