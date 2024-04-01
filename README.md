# go-load-balancer-example

simple load balancing in go implementing the Round Robin algorithm

## structure

![alt text](https://github.com/KaitoXCode/go-load-balancer-example/blob/master/public/flow-diagram.jpg?raw=true)

## make cmds

run all web services on ports: [':6969', ':6970', '6971']<br />
`make web-up`<br /> drop all web services on ports: [':6969', ':6970',
'6971']<br /> `make web-down`<br /> run web service on port: ':6969'<br />
`make web69-up`<br /> drop web service on port: ':6969'<br />
`make web69-down`<br /> run web service on port: ':6970'<br />
`make web70-up`<br /> drop web service on port: ':6970<br />
`make web70-down`<br /> run web service on port: ':6971'<br />
`make web71-up`<br /> drop web service on port: ':6971<br />
`make web71-down`<br /> run load balancer on port: ':7070'<br />
`make lb-up`<br /> drop load balancer on port: ':7070'<br />
`make lb-down`<br /> run all services:<br /> `make docker-up`<br /> stop all
services:<br /> `make docker-down`<br />

## usage

after starting all docker services, start pinging the load balancer and see how
it passes the request itteratively to each web service

![alt text](https://github.com/KaitoXCode/go-load-balancer-example/blob/master/public/usage.jpg?raw=true)
