FROM caddy:2
ADD public /
CMD caddy file-server --root public
