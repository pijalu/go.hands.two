FROM golang:1-alpine
RUN apk --no-cache add ca-certificates git
RUN go get -v github.com/golang/dep/cmd/dep
WORKDIR /go/src/github.com/pijalu/go.hands.two/frinsultgate
COPY . .
RUN dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fg .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /bin
COPY --from=0 /go/src/github.com/pijalu/go.hands.two/frinsultgate/fg .
ENTRYPOINT ["/bin/fg"]
