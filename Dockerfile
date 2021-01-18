###################################################
# Plain env as basement and for local development #
###################################################
FROM golang:alpine as env

MAINTAINER ayubirazeh@gmail.com

# Add support for Delve debugger.
RUN apk add --no-cache ca-certificates git
RUN go get github.com/derekparker/delve/cmd/dlv

##########################################################
# Prepare a build container with all dependencies inside #
##########################################################
FROM env as builder
WORKDIR $GOPATH/src/github.com/ayubirz/rent-vs-buy-grpc-go
COPY ./ .
RUN go build -o /go/bin/main cmd/main.go

###########################################
# Create clean container with binary only #
###########################################
FROM alpine as exec

RUN apk add --update bash ca-certificates

WORKDIR /app
COPY --from=builder /go/bin/main ./

ENTRYPOINT ["/app/main"]