FROM alpine:3.5
MAINTAINER Ce Gao(gaocegege) <gaocegege@hotmail.com>

RUN apk add --no-cache git \
    ruby \
    ruby-irb \
    ruby-rdoc \
    && gem install github_changelog_generator

COPY maintainer /bin

VOLUME /workdir
WORKDIR /workdir

ENTRYPOINT ["maintainer"]
