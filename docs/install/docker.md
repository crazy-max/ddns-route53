# Installation with Docker

## About

ddns-route53 provides automatically updated Docker :whale: images within several registries:

| Registry                                                                                         | Image                           |
|--------------------------------------------------------------------------------------------------|---------------------------------|
| [Docker Hub](https://hub.docker.com/r/crazymax/ddns-route53/)                             | `crazymax/ddns-route53`                 |
| [GitHub Container Registry](https://github.com/users/crazy-max/packages/container/package/ddns-route53)  | `ghcr.io/crazy-max/ddns-route53`        |

It is possible to always use the latest stable tag or to use another service that handles updating Docker images.

!!! note
    Want to be notified of new releases? Check out :bell: [Diun (Docker Image Update Notifier)](https://github.com/crazy-max/diun) project!

Following platforms for this image are available:

```shell
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

This reference setup guides users through the setup based on `docker-compose`, but the installation of `docker-compose`
is out of scope of this documentation. To install `docker-compose` itself, follow the official
[install instructions](https://docs.docker.com/compose/install/).

## Usage

```yaml
version: "3.5"

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

Edit this example with your preferences and run the following commands to bring up ddns-route53:

```shell
$ docker-compose up -d
$ docker-compose logs -f
```

Or use the following command:

```shell
$ docker run -d --name ddns-route53 \
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
$ docker-compose pull
$ docker-compose up -d
```

If you prefer to rely on the [`configuration file](../config/index.md#configuration-file) instead of
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

And your docker composition:

```yaml
version: "3.5"

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
