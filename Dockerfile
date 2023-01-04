FROM caddy:2
ADD public /srv
ADD Caddyfile /etc/caddy/Caddyfile
CMD caddy file-server
