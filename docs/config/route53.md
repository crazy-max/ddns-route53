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
```

### `name`

AWS Route 53 record set name.

!!! warning
    Don't forget to suffix with a dot

!!! example "Config file"
    ```yaml
    route53:
      recordsSet:
        - name: "ddns.example.com."
    ```

!!! abstract "Environment variables"
    * `DDNSR53_ROUTE53_RECORDSSET_<KEY>_NAME`

### `type`

AWS Route 53 record set type. Can be `A` or `AAAA`.

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
