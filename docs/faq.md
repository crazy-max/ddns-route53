# FAQ

## Timezone

By default, all interpretation and scheduling is done with your local timezone (`TZ` environment variable).

Cron schedule may also override the timezone to be interpreted in by providing an additional space-separated field
at the beginning of the cron spec, of the form `CRON_TZ=<timezone>`:

```shell
ddns-route53 --schedule "CRON_TZ=Asia/Tokyo */30 * * * *"
```

## Where is my public IP address sourced from?

The public IP address is retrieved from one of several providers. The first one to return a valid IP address is used.

### IPv4
* https://checkip.global.api.aws
* https://checkip.amazonaws.com
* https://cloudflare.com/cdn-cgi/trace
* https://ipv4.nsupdate.info/myip

### IPv6
* https://checkip.global.api.aws
* https://cloudflare.com/cdn-cgi/trace
* https://ipv6.nsupdate.info/myip
