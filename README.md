# What?
mockservice is a web service and client packed together in a single executable. 
The service listens for connections and the client is making requests to it at a 
constant rate.
The server responds back either correctly (http code 200) or wrongly (http code 500)in a statistically stable manner. I.e. the error rate is defined by user.
All these web transactions are captured by New Relic APM client and they are automatically 
sent to your NR account. There you can query them in insights (`FROM Transaction select...).

# How?
build: `go build`
run: `./mockservice`

build docker image: `docker build -t mockservice .`

# Why?
SRE teams build various applications around application metrics. Quite often the need arises to have a tightly controlled service that provides metrics that are stable, known beforehand and configurable. The team can use these metrics as control metrics to smoke test new monitoring applications or new features.

# REST config endpoint
By "GET"ting on `/set` endpoint and providing values for `err` and `rmp` you can set the error rate of the server (0-100) and the request per minute (0-?). 
Example: `http://localhost:8000/set/?err=10&rpm=30`
The settings will be applied on the fly.

# Arch
The app is very simple and minimal. Web server uses standard go libraries and the client is a simple goroutine that runs forever. A dockerfile can be used to build a slim docker image 
in case you want to deploy in a K8s cluster.
