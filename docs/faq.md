# FAQ

## Timezone

By default, scheduling uses your local timezone, controlled by the `TZ` environment variable.

You can also override the timezone in the cron expression by adding an extra space-separated field
at the beginning of the cron spec in the form `CRON_TZ=<timezone>`:

```shell
ddns-route53 --schedule "CRON_TZ=Asia/Tokyo */30 * * * *"
```

## Where is my public IP address sourced from?

The public IP address is retrieved from one of several providers. ddns-route53 uses the first one
that returns a valid IP address.

### IPv4
* https://checkip.global.api.aws
* https://checkip.amazonaws.com
* https://cloudflare.com/cdn-cgi/trace
* https://ipv4.nsupdate.info/myip

### IPv6
* https://checkip.global.api.aws
* https://cloudflare.com/cdn-cgi/trace
* https://ipv6.nsupdate.info/myip
