FROM caddy:2
ADD public/ /srv/public/
ADD Caddyfile /srv
CMD caddy run
