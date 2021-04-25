# Run as service on Debian based distro

## Using systemd

!!! warning
    Make sure to follow the instructions to [install from binary](binary.md) before.

To create a new service, paste this content in `/etc/systemd/system/ddns-route53.service`:

```
[Unit]
Description=ddns-route53
Documentation={{ config.site_url }}
After=syslog.target
After=network.target

[Service]
RestartSec=2s
Type=simple
User=ddns-route53
Group=ddns-route53
ExecStart=/usr/local/bin/ddns-route53 --config /etc/ddns-route53/ddns-route53.yml
Restart=always
#Environment=TZ=Europe/Paris
#Environment=AWS_ACCESS_KEY_ID=********
#Environment=AWS_SECRET_ACCESS_KEY=********
Environment=SCHEDULE=*/30 * * * *

[Install]
WantedBy=multi-user.target
```

Change the user, group, and other required startup values following your needs.

Enable and start ddns-route53 at boot:

```shell
sudo systemctl enable ddns-route53
sudo systemctl start ddns-route53
```

To view logs:

```shell
journalctl -fu ddns-route53.service
```
