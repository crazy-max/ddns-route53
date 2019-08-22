# Run as service on Debian based distro

## Using systemd

> :warning: Make sure to follow the instructions to [install from binary](binary.md) before.

Run the below command in a terminal:

```
sudo vim /etc/systemd/system/ddns-route53.service
```

Copy the sample [ddns-route53.service](https://github.com/crazy-max/ddns-route53/tree/master/.res/systemd/ddns-route53.service).

Change the user, group, and other required startup values following your needs.

Enable and start ddns-route53 at boot:

```
sudo systemctl enable ddns-route53
sudo systemctl start ddns-route53
```

To view logs:

```
journalctl -fu ddns-route53.service
```
