# LearnGo

[![Build Status](https://travis-ci.org/donvito/learngo.svg?branch=master)](https://travis-ci.org/donvito/learngo)

This repository contains code which I used to learn Go. Hoping it can be useful to someone, I've uploaded them here in github.

## learngo/helloworld

This is a very basic Go application which prints a simple string

## learngo/docker/api

This is a Go application which uses the Docker Go SDK and communicates with the Docker API. It outputs lists of containers, images, networks and swarm nodes.  Check my blog post about this https://www.melvinvivas.com/learning-go-with-docker-sdk/

## learngo/mongo-microservice

This is my first microservice using the Go Language which connects to a MongoDB NoSQL database, compiled using Docker a multi-stage build, and deployed as a Docker container. The microservice exposes 2 HTTP endpoints

* GET /jobs - this endpoint returns a list of jobs retrieved from a MongoDB database
* POST /jobs - this endpoint accepts a json string and saves it to a MongoDB database

Check my blog post about this
https://www.melvinvivas.com/my-first-go-microservice/
 
