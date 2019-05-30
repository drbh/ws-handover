# ws-handover

TLDR; Highly available websocket connections in Go  

[ws-handover demo video](https://raw.githubusercontent.com/drbh/ws-handover/master/images/screenshot.png)


## What is this?
Here we use Golang, Redis and Websockets to provide a simple HA solution to websockets. 

This is useful when we have a webapp that is reading from a database and serving that data over Websockets.

An example when this might be handy is when you are deploying a new update to a backend services that handles the websocket connections.

We pose a super simple solution - we run two instances of the back end and when a connection fails - we fallback onto the second applications address. 

This solution works since we hold all of the application sate outside of the server and in the DB or Redis instance in this case.

Follow the instructions to run this locally.


## Demo Video

[![ws-handover demo video](https://img.youtube.com/vi/xdy0oc9I_fg/0.jpg)](https://www.youtube.com/watch?v=xdy0oc9I_fg)


## Running

First start up a small Redis instance with Docker and expose the instance on the local `6379` port.
```
run --name some-redis -p 6379:6379 -d redis 
```

### Terminal A
```
go run updater.go 
```
This will connect to the Redis instance and increment a `counter` value. It +1 every second.


### Terminal B
```
go run service_a.go 
```

### Terminal C
```
go run service_b.go 
```

Now open up your browser to the `index.html` file. 