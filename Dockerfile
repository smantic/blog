FROM caddy:2
COPY ./public/ /
RUN caddy file-server
