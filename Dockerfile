FROM golang:1.26-alpine as builder

ARG CI_COMMIT_TAG

ENV GO111MODULE on

# Build project
WORKDIR /go/src/github.com/batazor/hcfc
COPY . .
RUN apk add --update git && \
  go get -u github.com/gobuffalo/packr/packr && \
  packr build cmd/hcfc/main.go && \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  go build \
  -a \
  -mod vendor \
  -ldflags "-w -s -X main.CI_COMMIT_TAG=$CI_COMMIT_TAG" \
  -installsuffix cgo \
  -o hcfc ./cmd/hcfc

FROM alpine

USER 10001

WORKDIR /app/
COPY --from=builder /go/src/github.com/batazor/hcfc/hcfc .
ENTRYPOINT ["./hcfc"]
