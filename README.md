# LearnGo

<!--- [![Build Status](https://travis-ci.org/donvito/learngo.svg?branch=master)](https://travis-ci.org/donvito/learngo) --->

This repository contains code which I used to learn Go. Hoping it can be useful to someone, I've uploaded them here in github.

## [learngo/helloworld](https://github.com/donvito/learngo/tree/master/helloworld)

This is a very basic Go application which prints a simple string

## [learngo/docker/api](https://github.com/donvito/learngo/docker/api)

This is a Go application which uses the Docker Go SDK and communicates with the Docker API. It outputs lists of containers, images, networks and swarm nodes.  Check my blog post about this https://www.melvinvivas.com/learning-go-with-docker-sdk/

## [learngo/mongo-microservice](https://github.com/donvito/learngo/tree/master/mongo-microservice)

This is my first microservice using the Go Language which connects to a MongoDB NoSQL database, compiled using Docker a multi-stage build, and deployed as a Docker container. The microservice exposes 2 HTTP endpoints

* GET /jobs - this endpoint returns a list of jobs retrieved from a MongoDB database
* POST /jobs - this endpoint accepts a json string and saves it to a MongoDB database

Check my blog post about this
https://www.melvinvivas.com/my-first-go-microservice/
 
## [learngo/rest-kafka-mongo-microservice](https://github.com/donvito/learngo/tree/master/rest-kafka-mongo-microservice)

In this example, I decoupled the saving of data to MongoDB and created another microservice to handle this. I also added Kafka to serve as the messaging layer so the microservices can work on its own concerns asynchrounously.

### Microservice 1
The REST microservice which receives data from a /POST http call to it. After receiving the request, it retrieves the data from the http request and saves it to Kafka. After saving, it responds to the caller with the same data sent via /POST

### Microservice 2
The microservice which subscribes to a topic in Kafka where Microservice 1 saves the data. Once a message is consumed by the microservice, it then saves the data to MongoDB.

Check my blog post about this
https://www.melvinvivas.com/developing-microservices-using-kafka-and-mongodb/

## [learngo/pointers](https://github.com/donvito/learngo/tree/master/pointers)

I've been studying about pointers in golang and these are the source codes which I used to understand pointers.
