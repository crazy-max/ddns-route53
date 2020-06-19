# Installing on a UniFi Security Gateway

## As a scheduled task (cron)

### On the UniFi Security Gateway

#### Install the binary

Determine which architecture the USG is using:

```
uname -a
```

Determine if the USG is using softfloat or hardfloat optimizations:

```
readelf -A /bin/sh | grep ABI_FP
```

> To ensure that ddns-route53 persists across reboots, it should be stored in the `/config/scripts` directory.

Install the ddns-route53 binary that matches the reported architecture & float type by following the [install from binary](binary.md) instructions.

#### Create the config file

Create the config file and place in your configuration:
```
vi /config/scripts/ddns-config.yml
```

### On the Cloud Key / Cloud Controller

#### Find the UniFi base path

By default, the paths are as follows:

* **UniFi Cloud Key:** `/srv/unifi`
* **Debian/Ubuntu Linux:** `/usr/lib/unifi`
* **Windows:** `%USERPROFILE%/Ubiquiti UniFi`
* **macOS:** `~/Library/Application Support/UniFi`

#### Create config.gateway.json

Create a file called config.gateway.json in your `<unifi_base>/data/sites/<site_id>` directory with the following contents:

> The `interval` field sets how often ddns-route53 will run (e.g. 1 minute).

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

See [UniFi - USG Advanced Configuration Using config.gateway.json](https://help.ui.com/hc/en-us/articles/215458888-UniFi-USG-Advanced-Configuration-Using-config-gateway-json) for more information.

#### Trigger a provision for the USG

See [How to Trigger a Provision](https://help.ui.com/hc/en-us/articles/360008240754#8) for instructions.
