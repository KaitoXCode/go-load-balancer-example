# go-load-balancer-example

simple load balancing implementation in go

# setup

required .env file envs:<br /> root directory (if running locally)<br /> e.g.
`home/$USER/projects/go-load-balancer-example/`<br /> port (if running single
web locally)<br /> e.g. `:6970`<br />

# make cmds

run all web services on ports: [':6969', ':6970', '6971']<br />
`make web-up`<br /> drop all web services on ports: [':6969', ':6970',
'6971']<br /> `make web-down`<br /> run web service on port: ':6969'<br />
`make web69-up`<br /> drop web service on port: ':6969'<br />
`make web69-down`<br /> run web service on port: ':6970'<br />
`make web70-up`<br /> drop web service on port: ':6970<br />
`make web70-down`<br /> run web service on port: ':6971'<br />
`make web71-up`<br /> drop web service on port: ':6971<br />
`make web71-down`<br />
