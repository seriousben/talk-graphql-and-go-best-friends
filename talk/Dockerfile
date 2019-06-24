FROM nginx:1.17-alpine

ARG REVEAL_JS_VERSION=3.8.0

WORKDIR /usr/share/nginx/html
RUN apk --update add curl && \
  mkdir slides && \
  mkdir revealjs && \
  curl -sSL https://github.com/hakimel/reveal.js/archive/${REVEAL_JS_VERSION}.tar.gz | tar xz --strip 1 -C ./revealjs
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY *.html slides/
COPY *.js slides/
COPY media slides/media