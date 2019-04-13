FROM golang:1.12-alpine as builder

ARG CI_COMMIT_TAG

# Build project
WORKDIR /go/src/github.com/batazor/hcfc
COPY . .
RUN CGO_ENABLED=0 \
  GOOS=linux \
  go build \
  -a \
  -ldflags "-w -s -X main.CI_COMMIT_TAG=$CI_COMMIT_TAG" \
  -installsuffix cgo \
  -o hcfc ./cmd/hcfc

FROM scratch

USER 10001

WORKDIR /app/
COPY --from=builder /go/src/github.com/batazor/hcfc/hcfc .
CMD ["./hcfc"]
