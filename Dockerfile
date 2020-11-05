# image used for the healthcheck binary
FROM golang:1.15-buster
COPY healthcheck/ /go/src/healthcheck/
RUN CGO_ENABLED=0 go build -ldflags '-w -s -extldflags "-static"' -o /healthcheck /go/src/healthcheck/
