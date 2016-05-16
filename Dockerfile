FROM golang
MAINTAINER travis.simon@nicta.com.au

# Copy across our src files
ADD . /go/src/github.com/travissimon/microservices/webserver
ADD webview.html /go/bin

# Build server
RUN go install github.com/travissimon/microservices/webserver

WORKDIR /go/bin

CMD webserver

# Listen on 8080
EXPOSE 8080