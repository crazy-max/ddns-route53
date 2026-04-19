# Configuration

## Overview

There are two ways to configure ddns-route53:

* In a [configuration file](#configuration-file)
* As [environment variables](#environment-variables)

Configuration sources are evaluated in the order listed above.

If an option is not set, its default value is used. The same applies to any nested option that is not set.

## Configuration file

At startup, ddns-route53 searches for a file named `ddns-route53.yml` (or `ddns-route53.yaml`) in:

* `/etc/ddns-route53/`
* `$XDG_CONFIG_HOME/`
* `$HOME/.config/`
* `.` _(the working directory)_

You can override this using the [`--config` flag or the `CONFIG` environment variable](../usage/cli.md).

??? example "ddns-route53.yml"
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

## Environment variables

All file-based configuration can be mapped to environment variables. For example, the following configuration:

??? example "ddns-route53.yml"
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

Can be expressed as:

??? example "environment variables"
    ```
    DDNSR53_CREDENTIALS_ACCESSKEYID=ABCDEFGHIJKLMNO123456
    DDNSR53_CREDENTIALS_SECRETACCESSKEY=abcdefgh123456IJKLMN+OPQRS7890+ABCDEFGH
    
    DDNSR53_ROUTE53_HOSTEDZONEID=ABCEEFG123456789
    DDNSR53_ROUTE53_RECORDSSET_0_NAME=ddns.example.com.
    DDNSR53_ROUTE53_RECORDSSET_0_TYPE=A
    DDNSR53_ROUTE53_RECORDSSET_0_TTL=300
    DDNSR53_ROUTE53_RECORDSSET_1_NAME=ddns.example.com.
    DDNSR53_ROUTE53_RECORDSSET_1_TYPE=AAAA
    DDNSR53_ROUTE53_RECORDSSET_1_TTL=300
    DDNSR53_ROUTE53_RECORDSSET_2_NAME=another.example2.com.
    DDNSR53_ROUTE53_RECORDSSET_2_TYPE=A
    DDNSR53_ROUTE53_RECORDSSET_2_TTL=600
    ```

## Reference

* [credentials](credentials.md)
* [route53](route53.md)
