FROM caddy:2
ADD public /srv
ADD Caddyfile /srv
CMD caddy run
