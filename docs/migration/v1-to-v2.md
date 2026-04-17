# ddns-route53 v1 to v2

## Configuration transposed into environment variables

All configuration can now be mapped to environment variables. Take a look at the
[documentation](../config/index.md#environment-variables) for more details.

`AWS_HOSTED_ZONE_ID` has been renamed to `DDNSR53_ROUTE53_HOSTEDZONEID` to follow the new environment variable naming scheme.

## All fields in configuration are now _camelCased_

To support environment variable mapping, all configuration fields are now _camelCased_:

* `credentials.access_key_id` > `credentials.accessKeyID`
* `credentials.secret_access_key` > `credentials.secretAccessKey`
* `route53.hosted_zone_id` > `route53.hostedZoneID`
* ...

??? example "v1"
    ```yaml
    credentials:
      access_key_id: "ABCDEFGHIJKLMNO123456"
      secret_access_key: "abcdefgh123456IJKLMN+OPQRS7890+ABCDEFGH"
    
    route53:
      hosted_zone_id: "ABCEEFG123456789"
      records_set:
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

??? example "v2"
    ```yaml
    credentials:
      accessKeyID: "ABCDEFGHIJKLMNO123456"
      secretAccessKey: "abcdefgh123456IJKLMN+OPQRS7890+ABCDEFGH"
    
    route53:
      hostedZoneID: "ABCEEFG123456789"
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
