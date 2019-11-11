# Installation from binary

## Download

ddns-route53 binaries are available in [releases](https://github.com/crazy-max/ddns-route53/releases) page.

Choose the archive matching the destination platform and extract ddns-route53:

```
wget -qO- https://github.com/crazy-max/ddns-route53/releases/download/v1.5.0/ddns-route53_1.5.0_linux_x86_64.tar.gz | tar -zxvf - ddns-route53
```

## Test

After getting the binary, it can be tested with `./ddns-route53 --help` or moved to a permanent location.

```
$ ./ddns-route53 --help
usage: ddns-route53 --config=CONFIG [<flags>]

Dynamic DNS for Amazon Route 53‎ on a time-based schedule. More info on
https://github.com/crazy-max/ddns-route53

Flags:
  --help               Show context-sensitive help (also try --help-long and
                       --help-man).
  --config=CONFIG      ddns-route53 configuration file.
  --schedule=SCHEDULE  CRON expression format.
  --timezone="UTC"     Timezone assigned to ddns-route53.
  --log-level="info"   Set log level.
  --log-json           Enable JSON logging output.
  --version            Show application version.
```

## Server configuration

Steps below are the recommended server configuration.

### Prepare environment

Create user to run ddns-route53 (ex. `ddnsr53`)

```
groupadd ddnsr53
useradd -s /bin/false -d /bin/null -g ddnsr53 ddnsr53
```

### Create required directory structure

```
mkdir /etc/ddns-route53
chown ddnsr53:ddnsr53 /etc/ddns-route53
chmod 770 /etc/ddns-route53
```

### Configuration

You must create your first [configuration](../configuration.md) file in `/etc/ddns-route53/ddns-route53.yml` and type:

```
chown ddnsr53:ddnsr53 /etc/ddns-route53/ddns-route53.yml
chmod 644 /etc/ddns-route53/ddns-route53.yml
```

### Copy binary to global location

```
cp ddns-route53 /usr/local/bin/ddns-route53
```

## Running ddns-route53

After the above steps, two options to run ddns-route53:

### 1. Creating a service file (recommended)

See how to create [Linux service](linux-service.md) to start ddns-route53 automatically.

### 2. Running from command-line/terminal

```
/usr/local/bin/ddns-route53 --config /etc/ddns-route53/ddns-route53.yml --schedule "0 */30 * * * *"
```

## Updating to a new version

You can update to a new version of ddns-route53 by stopping it, replacing the binary at `/usr/local/bin/ddns-route53` and restarting the instance.

If you have carried out the installation steps as described above, the binary should have the generic name `ddns-route53`. Do not change this, i.e. to include the version number.
