FROM golang:1.9.2 as builder
ARG SOURCE_LOCATION=/
WORKDIR ${SOURCE_LOCATION}
RUN go get -d -v github.com/gorilla/mux \
	&& go get -d -v gopkg.in/mgo.v2/bson \
	&& go get -d -v gopkg.in/mgo.v2
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
ARG SOURCE_LOCATION=/
RUN apk --no-cache add curl
EXPOSE 9090
WORKDIR /root/
COPY --from=builder ${SOURCE_LOCATION} .
CMD ["./app"]  
