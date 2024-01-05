# FAQ

## Timezone

By default, all interpretation and scheduling is done with your local timezone (`TZ` environment variable).

Cron schedule may also override the timezone to be interpreted in by providing an additional space-separated field
at the beginning of the cron spec, of the form `CRON_TZ=<timezone>`:

```shell
ddns-route53 --schedule "CRON_TZ=Asia/Tokyo */30 * * * *"
```

## Where is my public IP address sourced from?

Public IP address is sourced from one of several providers. The first to provide a valid IP address is used.

### IPv4
* https://ipv4.nsupdate.info/myip
* https://v4.ident.me
* https://ipv4.yunohost.org
* https://ipv4.wtfismyip.com/text

### IPv6
* https://ipv6.nsupdate.info/myip
* https://v6.ident.me
* https://ipv6.yunohost.org
* https://ipv6.wtfismyip.com/text
