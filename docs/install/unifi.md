# Installing on a UniFi Security Gateway

## Install the binary

Determine which architecture the USG is using:

```shell
$ uname -a
```

Determine if the USG is using `softfloat` or `hardfloat` optimizations:

```shell
$ readelf -A /bin/sh | grep ABI_FP
```

!!! tips
    To ensure that ddns-route53 persists across reboots, it should be stored in the `/config/scripts` directory.

Install the ddns-route53 binary that matches the reported architecture and float type by following
the [install from binary](binary.md) instructions.

Now you have to create your `ddns-route53.yml` [configuration file](../config/index.md#configuration-file) in
`/config/scripts/ddns-route53.yml`

## On the Cloud Key/Controller

### Find the UniFi base path

By default, the paths are as follows:

* **UniFi Cloud Key:** `/srv/unifi`
* **Debian/Ubuntu Linux:** `/usr/lib/unifi`
* **Windows:** `%USERPROFILE%/Ubiquiti UniFi`
* **macOS:** `~/Library/Application Support/UniFi`

### Create a `config.gateway.json`

Create a file called `config.gateway.json` in your `<unifi_base>/data/sites/<site_id>` directory with
the following contents:

```json
{
  "task-scheduler": {
    "task": {
      "dnsupdate": {
        "executable": {
          "path": "/config/scripts/ddns-route53 --config /config/scripts/ddns-config.yml"
        },
        "interval": "1m"
      }
    }
  }
}
```

!!! tips
    The `interval` field sets how often ddns-route53 will run (e.g. 1 minute).

!!! info
    See [UniFi - USG advanced configuration using `config.gateway.json`](https://help.ui.com/hc/en-us/articles/215458888-UniFi-USG-Advanced-Configuration-Using-config-gateway-json) for more information.

### Trigger a provision for the USG

See [How to trigger a provision](https://help.ui.com/hc/en-us/articles/360008240754#8) for instructions.
