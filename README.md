## About

This is a simple echo based webserver serving `/app/static` meant for lightweight proxy setups f.e. behind a haproxy.

## Usage

Sample Docker Compose file:
```
services:
  app:
    build:
      context: .
      dockerfile: .docker/server/Dockerfile
      args:
        - GOLANG_VARIANT=1.23
        - DEBIAN_VARIANT=bookworm
    restart: unless-stopped
    ports:
      - 127.0.0.1:8888:80
    volumes:
      - "/var/www/htdocs:/app/static"
```
