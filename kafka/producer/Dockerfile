FROM golang:1.9.2 as builder
ARG SOURCE_LOCATION=/
WORKDIR ${SOURCE_LOCATION}
RUN go get -d -v github.com/confluentinc/confluent-kafka-go/kafka
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build --ldflags '-extldflags "-static"' -a -installsuffix cgo -o app .

FROM alpine
RUN apk add --update --no-cache alpine-sdk bash python curl
WORKDIR /root
RUN git clone https://github.com/edenhill/librdkafka.git
WORKDIR /root/librdkafka
RUN /root/librdkafka/configure && make && make install
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ARG SOURCE_LOCATION=/
EXPOSE 9090
COPY --from=builder ${SOURCE_LOCATION} .
CMD ["./app"]
