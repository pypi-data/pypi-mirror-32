# EMR Launcher Consul

Consul plugin providing template functions for [http://www.github.com/tuneinc/emr_launcher](emr_launcher)

## Functions

* consul.get_value(key) - gets the value for a given Consul key

## ENV variables
* CONSUL_ACCESS_TOKEN - REQUIRED, sets the consul token
* CONSUL_HOST - REQUIRED, sets the consul host to use
* CONSUL_SCHEME - http or https default: https
* CONSUL_PORT - port consul api is running on default: 443


