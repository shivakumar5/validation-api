FROM go_base:latest

EXPOSE 9000

WORKDIR /srv

ENTRYPOINT ["/srv/bin/validation_api"]

ENV APP_NAME=validation_api \
    AUTH_ENV=development \
    HOME=/srv

# Binaries
RUN mkdir ./bin
COPY validation_api ./bin/validation_api

# Build as root; run as unprivileged user
USER www-data