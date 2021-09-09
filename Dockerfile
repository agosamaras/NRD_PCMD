# Build the server binary
FROM golang:1.16 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o server main.go

# Use distroless as minimal base image to package the server binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
# If I want a shell also
# FROM gcr.io/distroless/static:nonroot-debug as builder
WORKDIR /
COPY --from=builder /workspace/server .
USER nonroot:nonroot
EXPOSE 8080

ENTRYPOINT ["/server"]
