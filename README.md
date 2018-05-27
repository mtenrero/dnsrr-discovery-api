# dnsrr-discovery-api
DNS RoundRobin Discovery HTTP API for Docker Service Discovery within a Swarm Cluster using DNSRR.

It's written in pure Golang, so you're free to use any binary cause it's cross-platform compatible.

## Usage

This image exposes the port **9090** by default. It could be changed in a near future

## API

The whole API Definition can be found under the swagger directory in Swagger/OpenAPI format

### Basic Usage

**GET** api/_:hostname_

* Returns **200 OK** : List containing discovered IPs
```
["106.10.248.151","124.108.115.101","212.82.100.151","98.136.103.24","74.6.136.151"]
```

* Returns **204 No Content** : Hostname can't be discovered