# LearnGo

[![Build Status](https://travis-ci.org/donvito/learngo.svg?branch=master)](https://travis-ci.org/donvito/learngo)

This repository contains code which I used to learn Go. Hoping it can be useful to someone, I've uploaded them here in github.

## learngo/helloworld/ - contains a hello world application

Prints a simple string

## learngo/docker/ - contains Go applications communicating with the Docker SDK

Simple Go application which communicates with the Docker API. Sample lists containers, images, networks and swarm nodes.

## learngo/mongo-microservice/ - a simple microservice using the Go Language which connects to a MongoDB NoSQL database, compiled using Docker a multi-stage build, and deployed as a Docker container

The microservice exposes 2 HTTP endpoints

GET /jobs - this endpoint returns a list of jobs retrieved from a MongoDB database
POST /jobs - this endpoint accepts a json string and saves it to a MongoDB database
 
