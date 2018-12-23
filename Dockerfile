# Creates relayd daemon image

FROM golang:1.11.4-alpine
COPY . /go/src/github.com/rusenask/webhook-dispatcher
WORKDIR /go/src/github.com/rusenask/webhook-dispatcher
RUN apk add --no-cache git
RUN go install --ldflags="-s -w"

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/bin/webhook-dispatcher /webhook-dispatcher
ENTRYPOINT ["/webhook-dispatcher"]