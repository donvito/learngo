FROM golang:1.9.2 as builder
WORKDIR /Users/melvin/Documents/go/projects/src/github.com/donvito/learngo/connect_mongo/multistage
RUN go get -d -v github.com/gorilla/mux \
	&& go get -d -v gopkg.in/mgo.v2/bson \
	&& go get -d -v gopkg.in/mgo.v2
COPY connect-mongo-rest-sample.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add curl
EXPOSE 9090
WORKDIR /root/
COPY --from=builder /Users/melvin/Documents/go/projects/src/github.com/donvito/learngo/connect_mongo/multistage .
CMD ["./app"]  
