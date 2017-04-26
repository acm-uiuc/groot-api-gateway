FROM golang:1.8-alpine
MAINTAINER ACM@UIUC

# Get git
RUN apk add --update git bash

# Create app directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# Download and install external dependencies
RUN  go get -u github.com/golang/dep/... && \
     dep ensure
    
# Bundle app source
ADD . /usr/src/app

# Create symlink to GOPATH for groot
RUN mkdir -p $GOPATH/src/github.com/acm-uiuc && \
    ln -sf /usr/src/app $GOPATH/src/github.com/acm-uiuc/groot-api-gateway

# Create folder for client key database
RUN mkdir -p /var/groot-api-gateway/


# Build groot
ADD build.sh /usr/src/app
RUN ./build.sh

CMD ["./build/groot-api-gateway", "-u"]
