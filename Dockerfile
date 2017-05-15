
FROM golang:1.8-alpine
MAINTAINER ACM@UIUC

# Get git
RUN apk add --update git bash

# Bundle app source
ADD . $GOPATH/src/github.com/acm-uiuc/groot-api-gateway
WORKDIR $GOPATH/src/github.com/acm-uiuc/groot-api-gateway

# Download and install external dependencies
RUN  go get -u github.com/golang/dep/...
    
# Create folder for client key database
RUN mkdir -p /var/groot-api-gateway/

# Build groot
ADD build.sh $GOPATH/src/github.com/acm-uiuc/groot-api-gateway
RUN dep ensure && ./build.sh

CMD ["./build/groot-api-gateway"]