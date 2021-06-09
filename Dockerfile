FROM golang:1.15-alpine as builder

WORKDIR /devops-go-demo
# Copy the Go Modules manifests
COPY . .
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN  go mod download

RUN go build  -a -o devops-go-demo main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM alpine:3.12
WORKDIR /
COPY --from=builder /devops-go-demo/devops-go-demo .
ENTRYPOINT ["/devops-go-demo"]

