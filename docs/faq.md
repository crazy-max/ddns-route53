# FAQ

## Timezone

By default, all interpretation and scheduling is done with your local timezone (`TZ` environment variable).

Cron schedule may also override the timezone to be interpreted in by providing an additional space-separated field
at the beginning of the cron spec, of the form `CRON_TZ=<timezone>`:

```shell
$ ddns-route53 --schedule "CRON_TZ=Asia/Tokyo */30 * * * *"
```
