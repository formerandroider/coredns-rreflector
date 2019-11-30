# rreflector

## Name
_rreflector_ - Reflects the IP (with an optional host) back for all `in-addr.arpa` and `ip6.arpa` PTR requests handled by the plugin.

## Description

_rreflector_ handles queries for `.in-addr.arpa.` and `.ip6.arpa.` and responds with a PTR record consisting of the requested sub-name, reversed, with an optional HOST appended.

## Syntax

`rreflector [HOST]`

* HOST hostname to append to IP address

## Examples

Reflect the IP back for any rDNS request for an IP in the 192.168.0.0.24 range.

```
192.168.0.0/24 {
    rreflector
}
``` 

Reflect the IP back, appended with the `.example.com` root domain, for any rDNS request for an IP in the 192.168.0.0.24 range.

```
192.168.0.0/24 {
    rreflector example.com
}
``` 