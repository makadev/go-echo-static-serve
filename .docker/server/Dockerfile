ARG GOLANG_VARIANT=1.23
ARG DEBIAN_VARIANT=bookworm

## stage for golang build
FROM golang:${GOLANG_VARIANT}-${DEBIAN_VARIANT} AS golang-build

COPY cmd /app/cmd
COPY internal /app/internal
COPY go.mod /app/
COPY go.sum /app/

WORKDIR /app

COPY .docker/server/buildscripts /usr/local/bin/

RUN set -ex \
    && chmod +x /usr/local/bin/go-download-retry.sh \
    && /usr/local/bin/go-download-retry.sh \
    && go build -o bin/check-config cmd/config/main.go \
    && go build -o bin/_server cmd/web-server/main.go

## stage for running stuff
FROM debian:${DEBIAN_VARIANT} AS debian-run

WORKDIR /app

## setup default test user, sudo and install needed packages
ARG USERNAME=runner
ARG USER_UID=1000
ARG USER_GID=1000
RUN apt-get update \
    && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
    sudo \
    ## certificates just in case
    ca-certificates \
    ## good to have, other locale
    locales-all \
    && groupadd --gid ${USER_GID} ${USERNAME} \
    && useradd --uid ${USER_UID} --gid ${USER_GID} -m ${USERNAME} \
    && echo ${USERNAME} ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/${USERNAME} \
    && chmod 0440 /etc/sudoers.d/${USERNAME} \
    && apt-get autoremove -y && apt-get clean -y && rm -rf /var/lib/apt/lists/*

## copy server/cmd binaries
COPY --from=golang-build /app/bin /usr/local/bin/

## add runtime/helper scripts
COPY .docker/server/scripts /usr/local/bin/

RUN set -ex \
    && chmod +x /usr/local/bin/start-server.sh

COPY static /app/static

## chown app-static path
RUN set -ex \
    && chown -R ${USERNAME}:${USERNAME} /app/static

## add default config
COPY .docker/server/config/baseconfig.yaml /app/config.yaml

USER $USERNAME

WORKDIR /app

CMD ["/usr/local/bin/start-server.sh"]
