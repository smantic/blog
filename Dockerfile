FROM caddy:2
ADD public /srv
COPY Caddyfile /etc/caddy/Caddyfile
CMD caddy file-server
