FROM node:10-alpine
RUN npm install -g @angular/cli
WORKDIR /fif
COPY fif/ .
RUN npm install && ng build --prod

FROM golang:1-alpine
RUN apk --no-cache add ca-certificates git
WORKDIR /go/src/github.com/pijalu/go.hands.two/frinsultfront
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ff .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /fif/dist/fif
COPY --from=0 /fif/dist/fif/ .
WORKDIR /
COPY --from=1 /go/src/github.com/pijalu/go.hands.two/frinsultfront/ff .
ENTRYPOINT ["/ff"]
