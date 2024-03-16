FROM golang:alpine AS builder

RUN apk update && \
    apk add curl \
            git \
            bash \
            make \
            ca-certificates && \
    rm -rf /var/cache/apk/*

WORKDIR /app

# copy go module files first
COPY go.* ./
RUN go mod download
RUN go mod verify

# copy source files and build the binary
COPY . .
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
RUN mkdir -p /var/log/app
WORKDIR /app/

COPY --from=builder /app/server .
COPY --from=builder /app/cmd/server/entrypoint.sh .
COPY --from=builder /app/config/*.yml ./config/
RUN ls -la
ENTRYPOINT ["./entrypoint.sh"]