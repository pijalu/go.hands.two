FROM golang:1-alpine
RUN apk --no-cache add ca-certificates git gcc musl-dev
RUN go get -v github.com/golang/dep/cmd/dep
WORKDIR /go/src/github.com/pijalu/go.hands.two/frinsultdata
COPY . .
RUN dep ensure -v
RUN GOOS=linux go build -a -installsuffix cgo -o fd .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /bin
COPY --from=0 /go/src/github.com/pijalu/go.hands.two/frinsultdata/fd .
ENTRYPOINT ["/bin/fd"]
