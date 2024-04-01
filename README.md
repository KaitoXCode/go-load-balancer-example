# go-load-balancer-example

simple load balancing in go implementing the Round Robin algorithm

## structure

![alt text](https://github.com/KaitoXCode/go-load-balancer-example/blob/master/public/flow-diagram.jpg?raw=true)

## make cmds

run all services:<br /> `make docker-up`<br /> stop all services:<br />
`make docker-down`<br /> send ping to load balancer:<br /> `make ping`<br />

## usage

after starting all docker services, start pinging the load balancer and see how
it passes the request itteratively to each web service

![alt text](https://github.com/KaitoXCode/go-load-balancer-example/blob/master/public/usage.jpg?raw=true)
