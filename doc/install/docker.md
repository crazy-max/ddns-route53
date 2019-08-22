# Installation with Docker

ddns-route53 provides automatically updated Docker :whale: images within [Docker Hub](https://hub.docker.com/r/crazymax/ddns-route53) and [Quay](https://quay.io/repository/crazymax/ddns-route53). It is possible to always use the latest stable tag or to use another service that handles updating Docker images.

Environment variables can be used within your container :

* `TZ` : Timezone assigned to ddns-route53
* `SCHEDULE` : [CRON expression](https://godoc.org/github.com/crazy-max/cron#hdr-CRON_Expression_Format) to schedule ddns-route53
* `LOG_LEVEL` : Log level output (default `info`)
* `LOG_JSON`: Enable JSON logging output (default `false`)

Docker compose is the recommended way to run this image. Copy the content of folder [.res/compose](https://github.com/crazy-max/ddns-route53/tree/master/.res/compose) in `/opt/ddns-route53/` on your host for example. Edit the compose and config file with your preferences and run the following commands:

```
docker-compose up -d
docker-compose logs -f
```

Or use the following command :

```
docker run -d --name ddns-route53 \
  -e "TZ=Europe/Paris" \
  -e "SCHEDULE=0 */30 * * * *" \
  -e "LOG_LEVEL=info" \
  -e "LOG_JSON=false" \
  -v "$(pwd)/ddns-route53.yml:/ddns-route53.yml:ro" \
  crazymax/ddns-route53:latest
```

To upgrade your installation to the latest release:

```
docker-compose pull
docker-compose up -d
```
