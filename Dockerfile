# Stage 1. Build the binary
FROM golang:1.11

# add a non-privileged user myapp
RUN useradd -u 10001 myapp

RUN mkdir -p /go/src/github.com/bapiens/k3mst
ADD . /go/src/github.com/bapiens/k3mst
WORKDIR /go/src/github.com/bapiens/k3mst

# build the binary with go build
RUN CGO_ENABLED=0 go build \
	-o bin/k3mst github.com/bapiens/k3mst/cmd/k3mst

# Stage 2. Run the binary
FROM scratch

ENV PORT 8080
ENV DIAG_PORT 8585

COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=0 /etc/passwd /etc/passwd
USER myapp

COPY --from=0 /go/src/github.com/bapiens/k3mst/bin/k3mst /k3mst
EXPOSE $PORT

CMD ["/k3mst"]
