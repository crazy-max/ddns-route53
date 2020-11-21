# Configuration

## Overview

There are two different ways to define configuration in ddns-route53:

* In a [configuration file](#configuration-file)
* As [environment variables](#environment-variables)

These ways are evaluated in the order listed above.

If no value was provided for a given option, a default value applies. Moreover, if an option has sub-options, and any of these sub-options is not specified, a default value will apply as well.

## Configuration file

At startup, ddns-route53 searches for a file named `ddns-route53.yml` (or `ddns-route53.yaml`) in:

* `/etc/ddns-route53/`
* `$XDG_CONFIG_HOME/`
* `$HOME/.config/`
* `.` _(the working directory)_

You can override this using the [`--config` flag or `CONFIG` env var](../usage/cli.md).

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
    
    ipProvider: "identme"
    ```

## Environment variables

All configuration from file can be transposed into environment variables. As an example, the following configuration:

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

    ipProvider: "identme"
    ```

Can be transposed to:

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

    DDNSR53_IPPROVIDER=identme
    ```

# IP Provider

You can choose between two different service to retrieve your external IP:

- [IdentMe](https://ident.me) - (use the value `identme`)
- [Ipify](https://www.ipify.org/) - (use the value `ipify`)

## Reference

* [credentials](credentials.md)
* [route53](route53.md)
