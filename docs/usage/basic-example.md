# Basic example

This section shows a minimal way to run ddns-route53.

## Setup

!!! warning
    Make sure to follow the instructions to [install from binary](../install/binary.md) before.

First, create a [`ddns-route53.yml` configuration](../config/index.md) file like this:

```yaml
# ./ddns-route53.yml
credentials:
  accessKeyID: "ABCDEFGHIJKLMNO123456"
  secretAccessKey: "abcdefgh123456IJKLMN+OPQRS7890+ABCDEFGH"

route53:
  hostedZoneID: "ABCEEFG123456789"
  recordsSet:
    - name: "ddns.example.com."
      type: "A"
      ttl: 300
```

That's it. You can now start ddns-route53 with the following command:

```shell
ddns-route53 --config ./ddns-route53.yml --schedule "*/30 * * * *"
```
