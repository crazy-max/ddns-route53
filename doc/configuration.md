# Configuration

Here is a YAML structure example:

```yml
credentials:
  access_key_id: "<AWS_ACCESS_KEY_ID>"
  secret_access_key: "<AWS_SECRET_ACCESS_KEY>"

route53:
  hosted_zone_id: "<HOSTED_ZONE_ID>"
  records_set:
    - name: "ddns.example.com."
      type: "A"
      ttl: 300
    - name: "another.example2.com."
      type: "A"
      ttl: 300
```

> Replace `<AWS_ACCESS_KEY_ID>`, `<AWS_SECRET_ACCESS_KEY>` and `<HOSTED_ZONE_ID>` with the correponded values.

* `credentials`
  * `access_key_id`: AWS Access Key.
  * `secret_access_key`: AWS Secret Key.
* `route53`
  * `hosted_zone_id`: AWS Route53 hosted zone ID.
  * `records_set`: Map of records set.
    * `name`: AWS Route53 record set name (don't forget to add a dot at the end).
    * `type`: AWS Route53 record set type, can be `A` or `AAAA`.
    * `ttl`: AWS Route53 record TTL (time to live) in seconds.
