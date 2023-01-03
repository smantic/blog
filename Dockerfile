FROM caddy:2
ADD public .
RUN caddy file-server
