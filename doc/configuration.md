# Configuration

Here is a YAML structure example:

```yml
credentials:
  access_key_id: "<awd_key_id>"
  secret_access_key: "<aws_secret_key>"

route53:
  hosted_zone_id: "<hosted_zone_id>"
  records_set:
    - name: "ddns.example.com."
      type: "A"
      ttl: 300
    - name: "another.example2.com."
      type: "A"
      ttl: 300
```

> Replace `<awd_key_id>`, `<aws_secret_key>` and `<hosted_zone_id>` with the correponded values.

* `credentials`
  * `access_key_id`: AWS Access Key.
  * `secret_access_key`: AWS Secret Key.
* `route53`
  * `hosted_zone_id`: AWS Route53 hosted zone ID.
  * `records_set`: Slice of records set.
    * `name`: AWS Route53 record set name (don't forget to add a dot at the end).
    * `type`: AWS Route53 record set type, can be `A` or `AAAA`.
    * `ttl`: AWS Route53 record TTL (time to live) in seconds.

## Environment variables

You can also use environment variables and omit some configuration entries:

* `AWS_ACCESS_KEY_ID` overrides `credentials.access_key_id`
* `AWS_SECRET_ACCESS_KEY` overrides `credentials.secret_access_key`
* `AWS_HOSTED_ZONE_ID` overrides `route53.hosted_zone_id`

```yml
route53:
  records_set:
    - name: "ddns.example.com."
      type: "A"
      ttl: 300
    - name: "another.example2.com."
      type: "A"
      ttl: 300
```
