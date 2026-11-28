# ddns-route53 v2 to v3

## Hosted Zone configuration key has changed

To support multiple AWS Hosted Zones for a single set of credentials, the former `route53` mapping contents are now
expressed as entries in a [YAML Sequence](https://yaml.org/spec/1.2.2/#21-collections) named `hostedZones`.

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

??? example "v3"
    ```yaml
    credentials:
      accessKeyID: "ABCDEFGHIJKLMNO123456"
      secretAccessKey: "abcdefgh123456IJKLMN+OPQRS7890+ABCDEFGH"

    hostedZones:
      - hostedZoneID: "ABCEEFG123456789"
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
