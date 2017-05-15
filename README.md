# GROOT

[![Build Status](https://travis-ci.org/acm-uiuc/groot-api-gateway.svg?branch=master)](https://travis-ci.org/acm-uiuc/groot-api-gateway)

This repo contains the API Gateway written in Fall of 2015
It provides the following capabilities:
  * Easy registration of services
  * Universal Authentication for the entire application - via an external authentication provided (Atlassian crowd)
  * Proxying API calls
  * Managing inter-service communication

When registering as service specify the data encoding and when requesting a resource though groot make the request using json.

Groot core development:

[![Join the chat at https://gitter.im/acm-uiuc/groot-development](https://badges.gitter.im/acm-uiuc/groot-development.svg)](https://gitter.im/acm-uiuc/groot-development?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Questions on how to add your app to Groot or use the Groot API:

[![Join the chat at https://gitter.im/acm-uiuc/groot-users](https://badges.gitter.im/acm-uiuc/groot-users.svg)](https://gitter.im/acm-uiuc/groot-users?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)



## RUNNING GROOT

Add the API spec in a new file (ex. todo.go) in the services package

There is a set of proxy api calls defined by the Arbor framework which the API Gateway is written in. 

AS OF 4/24/17
```go
/**
 *  Pass the http Request from the client and the ResponseWriter it expects
 *  Pass the target url of the backend service (not the url the client called)
 *  Pass the format of the service
 *  Pass a authorization token (optional)
 *  Will call the service and return the result to the client.
 **/
 arbor.GET(w http.ResponseWriter, url string, format string, token string, r *http.Request)
```
```go
 /**
  *  Pass the http Request from the client and the ResponseWriter it expects
  *  Pass the target url of the backend service (not the url the client called)
  *  Passes the encoded json(only format currently supported) to the service.
  *  Pass a authorization token (optional)
  *  Will call the service and return the result to the client.
  **/
  arbor.POST(w http.ResponseWriter, url string, format string, token string, r *http.Request)
```
```go
 /**
  *  Pass the http Request from the client and the ResponseWriter it expects
  *  Pass the target url of the backend service (not the url the client called)
  *  Passes the encoded json(only format currently supported) to the service.
  *  Pass a authorization token (optional)
  *  Will call the service and return the result to the client.
  **/
  arbor.PUT(w http.ResponseWriter, url string, format string, token string, r *http.Request)
```
```go
/**
 *  Pass the http Request from the client and the ResponseWriter it expects
 *  Pass the target url of the backend service (not the url the client called)
 *  Pass a authorization token (optional)
 *  Will call the service and return the result to the client.
 **/
 arbor.DELETE(w http.ResponseWriter, url string, format string, token string, r *http.Request)
```

All secret data should be kept in a file called config.go in the config directory

Make sure you have go dependencey management installed 

```
go get github.com/golang/dep/...

Install Dependencies [First time setup]

```sh
dep ensure 
```

Compile the service 

```sh
./build.sh
```

## CLI 
```sh
groot-api-gateway [-r | --register-client client_name] [-c | --check-registration token] [-u | --unsecured]
```

-r | --register-client *client_name*
> registers a client, generates a token

-c | --check-registration *token*
> checks if a token is valid and returns name of client

-u | --unsecured
> runs groot without the security layer 

*without args* 
> runs groot with the security layer


## Docker

#### Creating a client key:
```
docker-compose run groot-api-gateway ./build/groot-api-gateway -r <client_name>
```

#### Checking a client key: 
```
docker-compose run groot-api-gateway ./build/groot-api-gateway -c <client_key>
```

#### Config Values
To allow all services to work with the docker network, make sure to set these values in `config.go`:
```go
//Service hosts
const CreditsURL = "http://groot-credits-service:8765"
const EventsURL = "http://groot-events-service:8002"
const GigsURL = "http://groot-gigs-service:8964"
const GroupsURL = "http://groot-groups-service:9001"
const HardwareURL = "http://groot:4523/api/v1.0"
const MemesURL = "http://groot-meme-service:42069"
const MerchURL = "http://groot-merch-service:6969"
const QuotesURL = "http://groot-quotes-service:9494"
const RecruiterURL = "http://groot-recruiters-service:4567"
const UsersURL = "http://groot-users-service:8001"

//Arbor configurations
func LoadArborConfig() {
    security.AccessLogLocation = "/var/groot-api-gateway/access.log"
    security.ClientRegistryLocation = "/var/groot-api-gateway/clients.db"
    proxy.AccessControlPolicy = "*"
}
```

## License

This project is licensed under the University of Illinois/NCSA Open Source License. For a full copy of this license take a look at the LICENSE file. 

When contributing new files to this project, preappend the following header to the file as a comment: 

```
Copyright © 2017, ACM@UIUC

This file is part of the Groot Project.  
 
The Groot Project is open source software, released under the University of Illinois/NCSA Open Source License. 
You should have received a copy of this license in a file with the distribution.
```
