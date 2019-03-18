FROM golang:1.12-alpine as builder

# Build project
WORKDIR /go/src/github.com/batazor/hcfc
COPY . .
RUN CGO_ENABLED=0 GOOS=linux \
  go build \
  -a \
  -ldflags "-X main.CI_COMMIT_TAG=$CI_COMMIT_TAG" \
  -installsuffix cgo -o hcfc ./cmd/main

FROM alpine:latest

RUN addgroup -S 997 && adduser -S -g 997 997
USER 997

WORKDIR /app/
COPY --from=builder /go/src/github.com/batazor/hcfc .
CMD ["./hcfc"]
