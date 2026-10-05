# Route 53 configuration

## `hostedZoneID`

AWS Route 53 hosted zone ID.

!!! example "Config file"
    ```yaml
    route53:
      hostedZoneID: "ABCEEFG123456789"
    ```

!!! abstract "Environment variables"
    * `DDNSR53_ROUTE53_HOSTEDZONEID`

## `recordsSet`

List of record sets.

```yaml
route53:
  recordsSet:
    - name: "ddns.example.com."
      type: "A"
      ttl: 300
    - name: "ddns.example.com."
      type: "AAAA"
      ttl: 300
    - name: "another.example2.com."
      type: "A"
      ttl: 600
    - name: "_ssh._tcp.ddns.example.com."
      type: "SRV"
      ttl: 300
      priority: 0
      weight: 0
      port: 2032
      target: "ddns.example.com."
```

### `name`

AWS Route 53 record set name.

!!! warning
    Remember to end the record name with a trailing dot.

!!! example "Config file"
    ```yaml
    route53:
      recordsSet:
        - name: "ddns.example.com."
    ```

!!! abstract "Environment variables"
    * `DDNSR53_ROUTE53_RECORDSSET_<KEY>_NAME`

### `type`

AWS Route 53 record set type. Can be `A`, `AAAA`, or `SRV`.

SRV records use the configured priority, weight, port, and target rather than the
WAN IP address. They can be used on their own or alongside A and AAAA records.

!!! example "Config file"
    ```yaml
    route53:
      recordsSet:
        - name: "ddns.example.com."
          type: A
    ```

!!! abstract "Environment variables"
    * `DDNSR53_ROUTE53_RECORDSSET_<KEY>_TYPE`

### `ttl`

AWS Route 53 record TTL (time to live) in seconds.

!!! example "Config file"
    ```yaml
    route53:
      recordsSet:
        - name: "ddns.example.com."
          ttl: 300
    ```

!!! abstract "Environment variables"
    * `DDNSR53_ROUTE53_RECORDSSET_<KEY>_TTL`

### `priority`

SRV record priority, from `0` to `65535`. Defaults to `0`. Lower values are preferred.

!!! abstract "Environment variables"
    * `DDNSR53_ROUTE53_RECORDSSET_<KEY>_PRIORITY`

### `weight`

SRV record weight, from `0` to `65535`. Defaults to `0`. Controls the relative
selection weight among records with the same priority.

!!! abstract "Environment variables"
    * `DDNSR53_ROUTE53_RECORDSSET_<KEY>_WEIGHT`

### `port`

Service port, from `1` to `65535`. Required for SRV records.

!!! abstract "Environment variables"
    * `DDNSR53_ROUTE53_RECORDSSET_<KEY>_PORT`

### `target`

Target hostname. Required for SRV records. Use a fully qualified domain name with
a trailing dot, such as `ddns.example.com.`. This hostname should have an A or
AAAA record; the target is not replaced with the WAN IP address.

!!! abstract "Environment variables"
    * `DDNSR53_ROUTE53_RECORDSSET_<KEY>_TARGET`
