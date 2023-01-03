FROM caddy:2
ADD public /srv
CMD caddy file-server
