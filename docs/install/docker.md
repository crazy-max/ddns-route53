# Installation with Docker

## About

ddns-route53 publishes automatically updated container images to several registries:

| Registry                                                                                                | Image                            |
|---------------------------------------------------------------------------------------------------------|----------------------------------|
| [Docker Hub](https://hub.docker.com/r/crazymax/ddns-route53/)                                           | `crazymax/ddns-route53`          |
| [GitHub Container Registry](https://github.com/users/crazy-max/packages/container/package/ddns-route53) | `ghcr.io/crazy-max/ddns-route53` |

You can follow the latest stable tag, or use another service to automate Docker image updates.

!!! note
    Want to be notified of new releases? See [Diun (Docker Image Update Notifier)](https://github.com/crazy-max/diun).

The following platforms are available for this image:

```
$ docker run --rm mplatform/mquery crazymax/ddns-route53:latest
Image: crazymax/ddns-route53:latest
 * Manifest List: Yes
 * Supported platforms:
   - linux/amd64
   - linux/arm/v6
   - linux/arm/v7
   - linux/arm64
   - linux/386
   - linux/ppc64le
```

This reference setup uses `docker-compose`, but installing `docker-compose`
is out of scope for this documentation. To install it, follow the official
[install instructions](https://docs.docker.com/compose/install/).

## Usage

```yaml
services:
  ddns-route53:
    image: crazymax/ddns-route53:latest
    container_name: ddns-route53
    environment:
      - "TZ=Europe/Paris"
      - "SCHEDULE=*/30 * * * *"
      - "LOG_LEVEL=info"
      - "LOG_JSON=false"
      - "DDNSR53_CREDENTIALS_ACCESSKEYID=ABCDEFGHIJKLMNO123456"
      - "DDNSR53_CREDENTIALS_SECRETACCESSKEY=abcdefgh123456IJKLMN+OPQRS7890+ABCDEFGH"
      - "DDNSR53_ROUTE53_HOSTEDZONEID=ABCEEFG123456789"
      - "DDNSR53_ROUTE53_RECORDSSET_0_NAME=ddns.example.com."
      - "DDNSR53_ROUTE53_RECORDSSET_0_TYPE=A"
      - "DDNSR53_ROUTE53_RECORDSSET_0_TTL=300"
    restart: always
```

Adjust this example to match your environment, then run the following commands to start ddns-route53:

```shell
docker-compose up -d
docker-compose logs -f
```

Or run it directly:

```shell
docker run -d --name ddns-route53 \
  -e "TZ=Europe/Paris" \
  -e "SCHEDULE=*/30 * * * *" \
  -e "LOG_LEVEL=info" \
  -e "LOG_JSON=false" \
  -e "DDNSR53_CREDENTIALS_ACCESSKEYID=ABCDEFGHIJKLMNO123456" \
  -e "DDNSR53_CREDENTIALS_SECRETACCESSKEY=abcdefgh123456IJKLMN+OPQRS7890+ABCDEFGH" \
  -e "DDNSR53_ROUTE53_HOSTEDZONEID=ABCEEFG123456789" \
  -e "DDNSR53_ROUTE53_RECORDSSET_0_NAME=ddns.example.com." \
  -e "DDNSR53_ROUTE53_RECORDSSET_0_TYPE=A" \
  -e "DDNSR53_ROUTE53_RECORDSSET_0_TTL=300" \
  crazymax/ddns-route53:latest
```

To upgrade your installation to the latest release:

```shell
docker-compose pull
docker-compose up -d
```

If you prefer to rely on the [`configuration file`](../config/index.md#configuration-file) instead of
environment variables:

```yaml
# ./ddns-route53.yml
credentials:
  accessKeyID: "ABCDEFGHIJKLMNO123456"
  secretAccessKey: "abcdefgh123456IJKLMN+OPQRS7890+ABCDEFGH"

route53:
  hostedZoneID: "ABCEEFG123456789"
  recordsSet:
    - name: "ddns.example.com."
      type: "A"
      ttl: 300
```

Then update your Compose configuration:

```yaml
services:
  ddns-route53:
    image: crazymax/ddns-route53:latest
    container_name: ddns-route53
    volumes:
      - "./ddns-route53.yml:/ddns-route53.yml:ro"
    environment:
      - "TZ=Europe/Paris"
      - "SCHEDULE=*/30 * * * *"
      - "LOG_LEVEL=info"
      - "LOG_JSON=false"
    restart: always
```
