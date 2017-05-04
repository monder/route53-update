## route53-update

[![Go Report Card](https://goreportcard.com/badge/github.com/monder/route53-update)](https://goreportcard.com/report/github.com/monder/route53-update)
[![license](https://img.shields.io/github/license/monder/route53-update.svg?maxAge=2592000&style=flat-square)]()
[![GitHub tag](https://img.shields.io/github/tag/monder/route53-update.svg?style=flat-square)]()

Small utility to quickly update domains in route53 zone

### Running

```
./route53-update ZONE_ID DOMAIN IP [TTL] 

```
* `ZONE_ID` - The id of route53 hosted zone
* `DOMAIN` - Fully qualified domain name to update
* `IP` - Value for domains A-record
* `TTL` - TTL for the record (default: `1`)

Silently exits with code `0` when successful

### Example

```
./route53-update ZXSVKTZADI4VW www.example.com 127.0.0.1 300
```
