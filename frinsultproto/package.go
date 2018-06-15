package frinsultproto

/*
 Protobuf needs to be installed
 eg: brew install protobuf
*/

// Go generate items
//go:generate go get -v github.com/golang/protobuf/protoc-gen-go
//go:generate go get -v github.com/micro/protoc-gen-micro
//go:generate protoc -I$GOPATH/src --go_out=$GOPATH/src --micro_out=$GOPATH/src $PWD/frinsult.proto
