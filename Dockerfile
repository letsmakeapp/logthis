FROM golang:1.23.1-alpine3.20 AS builder

ENV CGO_ENABLED=0
COPY . .
RUN go mod download && go build -o /app/main .

FROM alpine:3.20 AS runner
COPY --from=builder /app/main /app/main
CMD [ "/app/main" ]