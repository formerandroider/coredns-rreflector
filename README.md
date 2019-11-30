# RReflector

## Named
_rreflector_ - Reflects the IP (with an optional host) back for all `in-addr.arpa` and `ip6.arpa` PTR requests handled by the plugin.

## Description

_rreflector_ handles queries for `.in-addr.arpa.` and `.ip6.arpa.` and responds with a PTR record consisting of the requested sub-name, reversed, with an optional HOST appended.

## Syntax

_rreflector_ `[host]`

## Examples

```
192.168.0.0/24 {
    rreflector
}
``` 

```
192.168.0.0/24 {
    rreflector example.com
}
``` 