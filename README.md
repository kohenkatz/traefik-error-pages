# traefik-error-pages

A simple Go server for [Traefik](https://doc.traefik.io/traefik/)'s [ErrorPage](https://doc.traefik.io/traefik/middlewares/errorpages/) middleware.

Page design and content based on https://github.com/HttpErrorPages/HttpErrorPages.

Inspired by similar work by [Guillaume Briday](https://github.com/guillaumebriday/traefik-custom-error-pages)
and [Prasad Tengse](https://github.com/tprasadtp/traefik-error-pages).

## Setup Instructions

### Set up a container running this project in your `docker-compose.yml`

```yml
errorpage:
  restart: unless-stopped
  image: kohenkatz/traefik-error-pages:latest
  labels:
    - traefik.enable=true
    - traefik.http.routers.errorpage.entrypoints=websecure
    - "traefik.http.routers.errorpage.rule=HostRegexp(`{host:.+}`)"
    - traefik.http.services.errorpage.loadbalancer.server.port=3000
    - traefik.http.middlewares.errorpage.errors.status=400-599
    - traefik.http.middlewares.errorpage.errors.service=errorpage
    - traefik.http.middlewares.errorpage.errors.query=/HTTP{status}.html
    # Don't show the root file list
    - "traefik.http.middlewares.no-index.replacepathregex.regex=^/$$"
    - "traefik.http.middlewares.no-index.replacepathregex.replacement=/HTTP404.html"
    - traefik.http.routers.errorpage.middlewares=no-index
```

### Tell your other containers to use this one for error pages

#### Individually per-service

```yml
- traefik.http.routers.myapplication.middlewares=errorpage
```

Use commas to separate multiple middlewares:

```yml
- traefik.http.routers.myapplication.middlewares=securityHeaders@file,errorpage
```

#### Globally per entrypoint

```toml
[entryPoints]
  [entryPoints.websecure]
    address = ":443"
    [entryPoints.websecure.http]
      middlewares = ["errorpage@docker"]
```

## Build Instructions

This project is published to [Docker Hub](https://hub.docker.com/r/kohenkatz/traefik-error-pages) and
[GitHub Containers](https://github.com/kohenkatz/traefik-error-pages/pkgs/container/traefik-error-pages).

```sh
docker buildx build \
    --platform linux/amd64,linux/arm64 \
    --tag kohenkatz/traefik-error-pages:latest \
    --tag ghcr.io/kohenkatz/traefik-error-pages:latest \
    --push \
    .
```

In order to use the command above, you must be logged into Docker Hub and GitHub Containers.
