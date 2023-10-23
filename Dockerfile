FROM golang:1.21-alpine AS builder
WORKDIR /eff-mobile
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -v -o effmobile ./cmd/main.go

FROM alpine:3.16
COPY --from=builder /eff-mobile/effmobile /
USER nobody
CMD ["/effmobile"]
