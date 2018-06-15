# Go Hands Two
The second gomming of the go workshop

## env
A env helper package

## postman
A Postman json to help with dev

## frinsultfront
That's the front end !
Nothing to do here... just a angular6 front served by a simple go app

## frinsultproto
This project contains the protobuf service definitions and message models.
The source package.go contains the go:generate commands to (re)generate the go files.
### Notes ###
* You need to have protobuf compiler installed - see https://github.com/google/protobuf
* proto plugins are installed automatically with go generate

## frinsultgate
This is the gateway service. It's receive REST request and will pass them to frinsultdata, using go-micro framework

### Uses
* https://github.com/gorilla/mux to route http requests (server)
* https://github.com/micro/go-micro as service framework(client)

### TODOs
You will need to implement downvote method !
Upvote may be of some help

## frinsultdata
This is our data service. It offers client storage

### Uses
* http://gorm.io/ as ORM framework
* https://github.com/micro/go-micro as service framework(server)

### TODOs
Complete GetFrinsults in repository.go

## k8s
Kubernetes deploys

