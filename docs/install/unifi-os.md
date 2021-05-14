# Installing on UniFi OS-based devices

This guide is for devices running UniFi OS such as:

* UDM
* UDM Pro

## Install on-boot-script

Follow the [guide](https://github.com/boostchicken/udm-utilities/tree/master/on-boot-script#steps).

## Create the config file

Create the config directory.

```shell
mkdir /mnt/data/ddns-route53
```

Create your `ddns-route53.yml` [configuration file](../config/index.md#configuration-file) and save
it as `/mnt/data/ddns-route53/ddns-route53.yml`.

Fix the permissions of the config file.

```shell
chown 1000:1000 /mnt/data/ddns-route53/ddns-route53.yml
```

## Setup the service

Copy the following configuration and save it to `/mnt/data/on_boot.d/20-ddns-route53.sh`. Be sure
to change the `TZ=America/Vancouver` line to your own timezone.

```shell
#!/bin/sh

CNI_PATH=/mnt/data/podman/cni
if [ ! -f "$CNI_PATH"/macvlan ]; then
    mkdir -p $CNI_PATH
    curl -L https://github.com/containernetworking/plugins/releases/download/v0.8.6/cni-plugins-linux-arm64-v0.8.6.tgz | tar -xz -C $CNI_PATH
fi

mkdir -p /opt/cni
rm -f /opt/cni/bin
ln -s $CNI_PATH /opt/cni/bin

for file in "$CNI_PATH"/*.conflist
do
    if [ -f "$file" ]; then
        link_path="/etc/cni/net.d/$(basename "$file")"
        rm -f "$link_path"
        ln -s "$file" "$link_path"
    fi
done

podman run -d --restart always \
  --name ddns-route53 \
  --hostname ddns-route53 \
  -v "/mnt/data/ddns-route53/ddns-route53.yml:/ddns-route53.yml:ro" \
  -e "TZ=America/Vancouver" \
  -e "SCHEDULE=*/30 * * * *" \
  -e "LOG_LEVEL=info" \
  -e "LOG_JSON=false" \
  docker.io/crazymax/ddns-route53:latest
```

Make the script executable.

```shell
chmod +x /mnt/data/on_boot.d/20-ddns-route53.sh
```

## Start the service

Start the service.

```shell
/mnt/data/on_boot.d/20-ddns-route53.sh
```

The next time the device restarts the service will automatically start.

## Checking the logs

Check the logs with podman.

```shell
podman logs ddns-route53
```
