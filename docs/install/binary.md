# Installation from binary

## Download

ddns-route53 binaries are available on [releases]({{ config.repo_url }}releases/latest) page.

Choose the archive matching the destination platform:

* [`ddns-route53_{{ git.tag | trim('v') }}_darwin_arm64.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_darwin_arm64.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_darwin_x86_64.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_darwin_x86_64.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_freebsd_i386.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_freebsd_i386.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_freebsd_x86_64.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_freebsd_x86_64.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_linux_arm64.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_linux_arm64.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_linux_armv5.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_linux_armv5.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_linux_armv6.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_linux_armv6.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_linux_armv7.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_linux_armv7.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_linux_i386.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_linux_i386.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_linux_mips64le_hardfloat.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_linux_mips64le_hardfloat.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_linux_mips64le_softfloat.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_linux_mips64le_softfloat.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_linux_mips64_hardfloat.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_linux_mips64_hardfloat.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_linux_mips64_softfloat.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_linux_mips64_softfloat.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_linux_mipsle_hardfloat.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_linux_mipsle_hardfloat.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_linux_mipsle_softfloat.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_linux_mipsle_softfloat.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_linux_mips_hardfloat.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_linux_mips_hardfloat.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_linux_mips_softfloat.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_linux_mips_softfloat.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_linux_ppc64le.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_linux_ppc64le.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_linux_s390x.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_linux_s390x.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_linux_x86_64.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_linux_x86_64.tar.gz)
* [`ddns-route53_{{ git.tag | trim('v') }}_windows_i386.zip`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_windows_i386.zip)
* [`ddns-route53_{{ git.tag | trim('v') }}_windows_x86_64.zip`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_windows_x86_64.zip)

And extract ddns-route53:

```shell
wget -qO- {{ config.repo_url }}releases/download/v{{ git.tag | trim('v') }}/ddns-route53_{{ git.tag | trim('v') }}_linux_x86_64.tar.gz | tar -zxvf - ddns-route53
```

After getting the binary, it can be tested with [`./ddns-route53 --help`](../usage/cli.md) command and moved to a
permanent location.

## Server configuration

Steps below are the recommended server configuration.

### Prepare environment

Create user to run ddns-route53 (ex. `ddns-route53`)

```shell
groupadd ddns-route53
useradd -s /bin/false -d /bin/null -g ddns-route53 ddns-route53
```

### Create required directory structure

```shell
mkdir -p /var/lib/ddns-route53
chown ddns-route53:ddns-route53 /var/lib/ddns-route53/
chmod -R 750 /var/lib/ddns-route53/
mkdir /etc/ddns-route53
chown ddns-route53:ddns-route53 /etc/ddns-route53
chmod 770 /etc/ddns-route53
```

### Configuration

Create your first [configuration](../config/index.md) file in `/etc/ddns-route53/ddns-route53.yml` and type:

```shell
chown ddns-route53:ddns-route53 /etc/ddns-route53/ddns-route53.yml
chmod 644 /etc/ddns-route53/ddns-route53.yml
```

### Copy binary to global location

```shell
cp ddns-route53 /usr/local/bin/ddns-route53
```

## Running ddns-route53

After the above steps, two options to run ddns-route53:

### 1. Creating a service file (recommended)

See how to create [Linux service](linux-service.md) to start ddns-route53 automatically.

### 2. Running from terminal

```shell
/usr/local/bin/ddns-route53 \
  --config /etc/ddns-route53/ddns-route53.yml \
  --schedule "*/30 * * * *"
```

## Updating to a new version

You can update to a new version of ddns-route53 by stopping it, replacing the binary at `/usr/local/bin/ddns-route53`
and restarting the instance.

If you have carried out the installation steps as described above, the binary should have the generic name
`ddns-route53`. Do not change this, i.e. to include the version number.
