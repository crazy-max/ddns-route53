# WAN IP configuration

## `providers`

Optional WAN IP lookup providers.

If configured, the custom providers replace the built-in defaults for the corresponding address family.
Each endpoint must return a plain text IP address in the response body.

!!! example "Config file"
    ```yaml
    wanip:
      providers:
        ipv4:
          - "https://ipv4.example.com"
        ipv6:
          - "https://ipv6.example.com"
    ```

### `ipv4`

List of IPv4 provider URLs used for `A` record updates.

!!! abstract "Environment variables"
    * `DDNSR53_WANIP_PROVIDERS_IPV4`

### `ipv6`

List of IPv6 provider URLs used for `AAAA` record updates.

!!! abstract "Environment variables"
    * `DDNSR53_WANIP_PROVIDERS_IPV6`

!!! info
    Environment variable values use a comma-separated list, for example:
    `DDNSR53_WANIP_PROVIDERS_IPV4=https://ipv4.example.com,https://ipv4-backup.example.com`
