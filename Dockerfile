FROM golang:1.23.3-alpine AS builder
RUN apk add --no-cache gcc musl-dev
ENV CGO_ENABLED=1
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go \
	build \ 
	-trimpath \
	-o /bin/app \
	cmd/app/main.go

FROM alpine:3.20
COPY --from=builder /bin/app /bin/app
RUN mkdir -p /data
ENTRYPOINT ["/bin/app"]
