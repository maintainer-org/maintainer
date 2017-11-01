FROM alpine:3.5
LABEL maintainer="Ce Gao(gaocegege) <gaocegege@hotmail.com>"

RUN apk add --no-cache git \
    bash \
    wget \
    curl \
    ruby \
    ruby-irb \
    ruby-rdoc \
    && gem install github_changelog_generator \
    && wget https://raw.githubusercontent.com/ekalinin/github-markdown-toc/master/gh-md-toc \
    && chmod a+x gh-md-toc \
    && mv gh-md-toc /bin

# Build from scripts/build-for-alpine.sh
COPY maintainer /bin

VOLUME /workdir
WORKDIR /workdir

ENTRYPOINT ["maintainer"]
