FROM caddy:2
ADD public /
CMD caddy file-server
