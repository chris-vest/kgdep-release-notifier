FROM golang:1.12.4 as builder

ADD . /go/src/github.com/chris-vest/kgdep-release-notifier
WORKDIR /go/src/github.com/chris-vest/kgdep-release-notifier

ENV GO111MODULE=on

RUN make build

FROM alpine:3.6
RUN apk --no-cache add ca-certificates

COPY --from=builder /go/src/github.com/chris-vest/kgdep-release-notifier /bin/
ENTRYPOINT [ "/bin/kgdep-release-notifier" ]
