FROM golang:1.15.2 AS builder

WORKDIR /workspace
COPY . .
ENV CGO_ENABLED=0
RUN go build .

FROM scratch

WORKDIR /workspace

COPY --from=builder /workspace/hello-there ./hello-there

EXPOSE 8080

CMD ["/workspace/hello-there"]
