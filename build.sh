#!/bin/bash

mkdir -p log/prod

#########################################################################
#																		#
#																		#
#			I N S T A L L   G R O O T   A P I   G A T E W A Y 		    #
#																		#
#																		#
#########################################################################	
go get -u github.com/acm-uiuc/arbor/...

echo Building config
go install github.com/acm-uiuc/groot-api-gateway/config

echo Building services
go install github.com/acm-uiuc/groot-api-gateway/services

echo Building groot-api-gateway
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
mkdir -p build 
echo Placing binary in [PATH TO GROOT]/build
(cd $DIR; go build -o  $DIR/build/groot-api-gateway .) 
