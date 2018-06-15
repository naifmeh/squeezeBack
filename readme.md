# Backend for face recognition project

## Goals of the backend

We would like to introduce the following features in the project :

* Face authorization by Rest API (in Golang)
* RethinkDb usage for the NOSQL database
* Kafka as the multi-broken as a publish-subscripe pattern

Clients should be able to run on any platform, being a mobile ap, a website,
or an embedded system.

## Kafka

Kafka is message broker based on internal commit log, as it focus on storing massive amount of data on disk, allowing consumption in real-time or delayed, as the storing data is kept locally.

## REST API

### GoLang

Go is a recent language develloped by the one of the father of the C language. It is a compiled, garbage collected language whose syntax is very close to the C syntax, while maintaining a simplistic approach. Many principles have been avoided to provide simplicity of coding. 
Classes are replaced by analog to C, structures. Interfaces provide an abstraction of the object oriented interfaces while introducing differences with them.
GoLang has in mind multi-threading, by the use of goroutines and channel based communication. Goroutines does not involve parallelism but asynchronous execution. 
Those concepts makes it very easy to implement fast Go softwares.

Go is particularly known for the ease that comes with it when develloping web applications.
Some lines of code are enough to build a working web server and process requests.
Further discussion on go concepts will be available in this documentation.

#### The net/http library

